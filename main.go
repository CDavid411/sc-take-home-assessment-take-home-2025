package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID) //UUID

	res := folder.GetAllFolders() //[]Folder

	// example usage
	folderDriver := folder.NewDriver(res)              //IDriver interface
	orgFolder := folderDriver.GetFoldersByOrgID(orgID) //[]Folder

	folder.PrettyPrint(res)
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)
}
