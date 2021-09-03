package cic

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"runcic/containerimage/common"
	"runcic/utils"
	"strings"
	"syscall"
	"time"
)

type Runcic struct {
	ParentRootfs    *os.File
	RunWith         bool
	Name            string
	CicVolume       string
	ContainerID     string
	Images          []*common.Image
	Command         []string
	Envs            []string
	CopyEnv         bool
	Started         time.Time
	ImagePullPolicy ImagePullPolicy
}

func (r *Runcic) ImageArray() (imgs []string) {
	for i := 0; i < len(r.Images); i++ {
		imgs = append(imgs, r.Images[i].Image)
	}
	return
}
func (r *Runcic) Roorfs() (path string) {
	path = OverlayRoot + string(os.PathSeparator) + r.Name
	return
}

func (r *Runcic) mountops() string {
	lower := make([]string, 0)
	for i := 0; i < len(r.Images); i++ {
		lower = append(lower, r.Images[i].Lower...)
	}
	mountops := strings.Join([]string{
		"lowerdir=" + strings.Join(lower, ":"),
		"upperdir=" + r.CicVolume + string(os.PathSeparator) + "up",
		"workdir=" + r.CicVolume + string(os.PathSeparator) + "work",
	}, ",")
	return mountops
}

func (r *Runcic) rootfspath() (err error) {
	work, worke := os.Stat(r.Roorfs())
	if worke != nil {
		utils.Mkdirp(r.Roorfs())
	} else {

		if r.RunWith {
			if !work.IsDir() {
				err = errors.New("rootfs should be dir!" + r.Roorfs())
				logrus.Warnf(err.Error())
			}
			return
		} else {
			err = errors.New(fmt.Sprintf("%s rootfs %s already exist!", r.Name, r.Roorfs()))
			logrus.Errorf(err.Error())
			return
		}
	}
	logrus.Infof("rootfs ok,%s", r.Roorfs())
	return
}
func (r *Runcic) cicvolume() (err error) {
	if r.CicVolume == "" {
		err = errors.New("cicvolume not set!")
		logrus.Errorf(err.Error())
		return err
	}
	stat, fe := os.Stat(r.CicVolume)
	if fe != nil {
		logrus.Errorf(fe.Error())
		return fe
	}
	if !stat.IsDir() {
		err = errors.New("cicvolume should be dir!,fix your cicvolume: " + r.CicVolume)
		logrus.Errorf(err.Error())
		return
	}

	up, upe := os.Stat(r.CicVolume + "/" + "up")
	if upe != nil {
		err = os.Mkdir(r.CicVolume+"/"+"up", os.ModeDir)
		if err != nil {
			logrus.Errorf("mkdir cicvolume/up dir fail,error: %s", err.Error())
			return err
		}
	} else {
		if !up.IsDir() {
			err = errors.New("cicvolume/up should be dir!" + r.CicVolume)
			logrus.Warnf(err.Error())
		}
	}
	logrus.Infof("cicvolume updir ok,%s", r.CicVolume+"/"+"up")

	work, worke := os.Stat(r.CicVolume + "/" + "work")
	if worke != nil {
		err = os.Mkdir(r.CicVolume+"/"+"work", os.ModeDir)
		if err != nil {
			logrus.Errorf("mkdir cicvolume/work dir fail,error: %s", err.Error())
			return err
		}
	} else {
		if !work.IsDir() {
			err = errors.New("cicvolume/work should be dir!" + r.CicVolume)
			logrus.Warnf(err.Error())
		}
	}

	logrus.Infof("cicvolume workdir ok,%s", r.CicVolume+"/"+"work")
	return
}

//Execv
// https://github.com/opencontainers/runc/blob/master/libcontainer/system/linux.go
func Execv(cmd string, args []string, env []string) error {
	name, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}

	return Exec(name, args, env)
}

func (r *Runcic) mountoverlay() (err error) {
	mountops := r.mountops()
	err = syscall.Mount("overlay", r.Roorfs(), "overlay", 0, mountops)
	logrus.Infof("mount overlay overlay -o %s %s ", mountops, r.Roorfs())
	if err != nil {
		logrus.Errorf("mount overlay fail,errors %s", err.Error())
		return
	}

	return err
}
func realChroot(path string) (oldRootF *os.File, err error) {

	oldRootF, err = os.Open("/")

	logrus.Infof("chrooting %s", path)
	if err := syscall.Chroot(path); err != nil {
		return oldRootF, fmt.Errorf("Error after fallback to chroot: %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return oldRootF, fmt.Errorf("Error changing to new root after chroot: %v", err)
	}
	logrus.Infof("chroot success %s", path)
	return
}

func Exec(cmd string, args []string, env []string) error {
	for {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()
		err := syscall.Exec(cmd, args, env)
		if err != syscall.EINTR { //nolint:errorlint // unix errors are bare
			return err
		}
	}
}
