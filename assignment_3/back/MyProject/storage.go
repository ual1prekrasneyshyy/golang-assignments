package main

var items = []Item{
	{Id: 1, Name: "Xiaomi Redmi Note 13 Pro", Price: 120000, Quantity: 30, Rating: 4.9},
	{Id: 2, Name: "IPhone 16 Pro", Price: 800000, Quantity: 40, Rating: 4.95},
	{Id: 3, Name: "Samsung Galaxy S24", Price: 500000, Quantity: 35, Rating: 4.97},
	{Id: 4, Name: "MacBook Pro", Price: 1200000, Quantity: 20, Rating: 4.96},
	{Id: 5, Name: "Dell Vostro 3490", Price: 300000, Quantity: 25, Rating: 4.91},
}

var roles = []Role{
	{Id: 1, Name: "ROLE_ADMIN"},
	{Id: 2, Name: "ROLE_USER"},
}

var users = []User{
	{Id: 1, Username: "ualishka", Password: "qwerty", FullName: "Uali Amangaliyev", Age: 21, Roles: []Role{roles[0], roles[1]}},
	{Id: 2, Username: "customer_1", Password: "qwerty", FullName: "Customer 1", Age: 22, Roles: []Role{roles[1]}},
	{Id: 3, Username: "customer_2", Password: "qwerty", FullName: "Customer 2", Age: 20, Roles: []Role{roles[1]}},
}
