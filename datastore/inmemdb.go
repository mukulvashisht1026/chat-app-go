package datastore

var users []User
var groups []Group
var messages []Message
var groupsToUserMapping []GroupsToUserMapping


func AddUser(Id, userType, name string) {
	a := User{
		Id: Id,
		Type: userType,
		Name: name,
	}
	users = append(users, a)
}

func ViewAllUsers() {
	if len(users) == 0 {
		fmt.Println("currently there are no users present in the system")
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}  
}



