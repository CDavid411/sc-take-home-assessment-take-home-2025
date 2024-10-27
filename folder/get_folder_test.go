package folder_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()
	folders, err := folder.LoadSampleData("get_folder_sample.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	myDriver := folder.NewDriver(folders)

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{"alpha", uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), folders, []folder.Folder{
			{
				Name:  "bravo",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "alpha.bravo",
			},
			{
				Name:  "charlie",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "alpha.bravo.charlie",
			},
			{
				Name:  "delta",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "alpha.delta",
			},
		}}, {"bravo", uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), folders, []folder.Folder{
			{
				Name:  "charlie",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "alpha.bravo.charlie",
			},
		}}, {"charlie", uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), folders, []folder.Folder{}},
		{"echo", uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), folders, []folder.Folder{}},
		{"invalid_folder", uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), folders, nil},
		{"foxtrot", uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), folders, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myDriver.GetAllChildFolders(tt.orgID, tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nExpected: \n%v\nGot: \n%v\n", tt.want, got)
			}
		})
	}
}

//change OrgId "org1" as "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
//"org2" as "c1556e17-b7c0-45a3-a6ae-9546248fb18b"
