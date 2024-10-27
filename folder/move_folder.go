package folder

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	if name == dst {
		return nil, errors.New("cannot move a folder to itself")
	}

	folders := f.folders
	res := folders
	var fName Folder
	var fDst Folder
	for _, ff := range folders { //find target folders
		if ff.Name == name {
			fName = ff
		}
		if ff.Name == dst {
			fDst = ff
		}
	}

	if fName.Name == "" && fName.OrgId == uuid.Nil && fName.Paths == "" {
		return nil, errors.New("source folder does not exist")
	}
	if fDst.Name == "" && fDst.OrgId == uuid.Nil && fDst.Paths == "" {
		return nil, errors.New("destination folder does not exist")
	}
	if fName.OrgId != fDst.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}

	str := fDst.Paths
	str = str[:len(str)-len(dst)]
	for idx := strings.Index(str, name); idx != -1; idx = strings.Index(str, name) {
		if (idx == 0 || str[idx-1] == '.') && idx+len(name) <= len(str)-1 && str[idx+len(name)] == '.' {
			return nil, errors.New("cannot move a folder to a child of itself")
		}
		//find next dot
		str = str[idx:]
		idx = strings.Index(str, ".")
		if idx != -1 {
			str = str[idx:]
		}
	}

	//note folders changes here
	folders = f.GetAllChildFolders(fName.OrgId, name)
	folders = append(folders, fName)
	for i, ff := range folders {
		str := ff.Paths
		keep := str
		idx := strings.Index(str, name)
		for ; idx != -1; idx = strings.Index(str, name) { //find exact idex
			if (idx == 0 || str[idx-1] == '.') && idx+len(name) <= len(str)-1 && str[idx+len(name)] == '.' { //targetf.
				break
			} else if ff.Name == name { //fname.targetf
				idx = len(str) - len(name)
				break
			}

			//find next dot
			str = str[idx:]
			t := strings.Index(str, ".")
			if t != -1 {
				str = str[t:]
				idx += t
			} else {
				fmt.Println("IDriver.GetAllChildFolders() error")
				break
			}
		}
		//dst insert
		t := keep[idx:]
		folders[i].Paths = fDst.Paths + "." + t
	}

	//#Paths replacing
	//need IDriver.GetAllChildFolders() to keep folders order to run correctly
	j := 0
	joName := len(folders) - 1
	for i, ff := range res {
		if ff.Name == folders[j].Name {
			res[i].Paths = folders[j].Paths
			j++
		} else if ff.Name == folders[joName].Name {
			res[i].Paths = folders[joName].Paths
		}
	}

	return res, nil
}

/*
#instance to be considered
	fname.xxtargetfnn.targetf
	targetf
	fname.targetftargetf.targetf
	fname.targetf

*/
