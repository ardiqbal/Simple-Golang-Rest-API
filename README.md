# Simple-Golang-Rest-API

Simple REST API to manage class attendance using Go Language

## GET: Get All Classes
### Endpoint
address:port/classes
### Result Example
```
{
    "success": true,
    "message": [
        "OK"
    ],
    "data": {
        "classes": [
            {
                "id": "5891693d-5107-11e9-9624-",
                "class_name": "Software Testing",
                "class_time": "2019-02-10T21:00:00+0700",
                "room": "E4",
                "created_at": "2019-03-28T17:12:40+0700",
                "updated_at": "2019-03-28T17:12:40+0700"
            },
            {
                "id": "7a4b31ce-5127-11e9-9624-",
                "class_name": "Computer Network",
                "class_time": "2019-02-09T19:00:00+0700",
                "room": "E9",
                "created_at": "2019-03-28T21:02:40+0700",
                "updated_at": "2019-03-28T21:02:40+0700"
            }
        ]
    }
}
```

## GET: Get All Classes
### Endpoint
address:port/classes
### Result Example
```
{
    "success": true,
    "message": [
        "OK"
    ],
    "data": {
        "classes": [
            {
                "id": "5891693d-5107-11e9-9624-",
                "class_name": "Software Testing",
                "class_time": "2019-02-10T21:00:00+0700",
                "room": "E4",
                "created_at": "2019-03-28T17:12:40+0700",
                "updated_at": "2019-03-28T17:12:40+0700"
            },
            {
                "id": "7a4b31ce-5127-11e9-9624-",
                "class_name": "Computer Network",
                "class_time": "2019-02-09T19:00:00+0700",
                "room": "E9",
                "created_at": "2019-03-28T21:02:40+0700",
                "updated_at": "2019-03-28T21:02:40+0700"
            }
        ]
    }
}
```

## POST: Create New Class
### Endpoint
address:port/add-class
### Request Body
```
{
	"class_name":"Computer Network",
	"class_time":"2019-02-09 12:00:00",
	"room":"E9"
}
```

## POST: Add Student Attendance
### Endpoint
address:port/add-attendance
### Request Body
```
{
	"class_id":"7a4b31ce-5127-11e9-9624-",
	"student_id":"43bb63e2-5106-11e9-9624-"
}
```
