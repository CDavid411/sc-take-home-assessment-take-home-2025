package folder

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	folders := f.GetFoldersByOrgID(orgID)

	if len(folders) == 0 { //orgID doesn't exist
		fmt.Println("Error: Folder does not exist in the specified organization")
		return nil
	}

	flag := false
	for _, ff := range f.folders { //orgID isn't match with the folder
		if ff.Name == name {
			if ff.OrgId != orgID {
				fmt.Println("Error: Folder does not exist in the specified organization")
				return nil
			}
			flag = true
		}
	}

	if !flag {
		fmt.Println("Error: Folder does not exist")
		return nil
	}

	res := []Folder{}
	for _, f := range folders {
		str := f.Paths
		for idx := strings.Index(str, name); idx != -1; idx = strings.Index(str, name) {
			if (idx == 0 || str[idx-1] == '.') && idx+len(name) <= len(str)-1 && str[idx+len(name)] == '.' { //judge child
				res = append(res, f)
				break
			} else if f.Name == name { //we have this folder, but it's empty
				break
			}

			//find next dot
			str = str[idx:]
			idx = strings.Index(str, ".")
			if idx != -1 {
				str = str[idx:]
			}
		}
	}

	return res
}

/*
#instance to be considered
	#deny
	targetf
	fname.targetf
	fname.nntargetf.targetf
	fname.fntargetfnn.childf
	fname.fntargetf.childf
	fname.targetftargetf.targetf

	#accept
	targetf.childf
	fname.targetf.childf
	fname.xtargetfnn.targetf.childf
	fname.targetf.childf.targetf
*/
