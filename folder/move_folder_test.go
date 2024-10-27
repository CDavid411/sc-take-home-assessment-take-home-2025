package folder_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	// TODO: your tests here
	t.Parallel()
	folders, err := folder.LoadSampleData("move_folder_sample.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	myDriver := folder.NewDriver(folders)

	tests := [...]struct {
		name string
		dst  string
		want []folder.Folder
	}{
		// TODO: your tests here
		{"bravo", "delta", []folder.Folder{
			{
				Name:  "alpha",
				Paths: "alpha",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "bravo",
				Paths: "alpha.delta.bravo",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "charlie",
				Paths: "alpha.delta.bravo.charlie",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "delta",
				Paths: "alpha.delta",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "echo",
				Paths: "alpha.delta.echo",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "foxtrot",
				Paths: "foxtrot",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb18b"),
			},
			{
				Name:  "golf",
				Paths: "golf",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
		}}, {"bravo", "golf", []folder.Folder{
			{
				Name:  "alpha",
				Paths: "alpha",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "bravo",
				Paths: "golf.bravo",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "charlie",
				Paths: "golf.bravo.charlie",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "delta",
				Paths: "alpha.delta",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "echo",
				Paths: "alpha.delta.echo",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
			{
				Name:  "foxtrot",
				Paths: "foxtrot",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb18b"),
			},
			{
				Name:  "golf",
				Paths: "golf",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			},
		}}, {"bravo", "charlie", nil},
		{"bravo", "bravo", nil},
		{"bravo", "foxtrot", nil},
		{"invalid_folder", "delta", nil},
		{"bravo", "invalid_folder", nil},
	}
	for _, tt := range tests {
		t.Run(tt.dst, func(t *testing.T) {
			got, err := myDriver.MoveFolder(tt.name, tt.dst)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nExpected: \n%v\nGot: \n%v\n", tt.want, got)
			}
		})
	}
}

//change OrgId "org1" as "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
//"org2" as "c1556e17-b7c0-45a3-a6ae-9546248fb18b"
