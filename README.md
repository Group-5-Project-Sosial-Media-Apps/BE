# REST API of Social Media App Project

Berikut adalah list REST API yang dapat digunakan dalam proses testing untuk pembuatan Social Media App Project dengan menggunakan Postman / Insomnium.

Untuk detail request body dapat dilihat dalam OpenAPI di link berikut: https://app.swaggerhub.com/apis/SocialMediaApps/sosmed/1.0.0

# REST API

The REST API to the example app is described below.
Dibawah ini adalah list REST API meliputi Request dan Response.

## Login

### Request

`POST /login`

    localhost:8000/login

### Response

    {
	    "data": {
    		"user_id": 1,
    		"nama": "alta",
    		"username": "alta",
    		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE5OTExNDA0OTQsImlhdCI6as2wMTcwMTk5MTE0MDQ5MywiaWQiOjJ9.C7CF6Xey5aoLPFCAD2o83eZRZX6O4KOtT5vgIvNP-fU"
	            },
	    "message": "login success"
    }

## Register

### Request

`POST /register`

    localhost:8000/register

### Response

    {
    	"data": {
    		"nama": "alta",
    		"username": "alta12",
    		"email": "alta@alterra.id"
    	        },
    	"message": "success create data"
    }

## Get User by UserID

### Request

`GET /user/{user_id}`

    localhost:8000/user/id

### Response

    {
    	"data": {
    		"user_id": 1,
    		"nama": "agus",
    		"username": "alta",
    		"email": "alta@alterra.id",
    		"foto": "https://res.cloudinary.com/djoxsmzq4/image/upload/v1701991113/sosmed/jdgasy0dgpqcbankmpp9.jpg"
    	        },
    	"message": "get user by userID successful"
    }

## Update User by UserID

### Request

`PUT /user/{user_id}`

    localhost:8000/updateuser

### Response

    {
    	"data": {
    		"UserID": 2,
    		"Nama": "agusan",
    		"UserName": "alta",
    		"Foto": "https://res.cloudinary.com/djoxsmzq4/image/upload/v1701991156/sosmed/njv4i0jvksdlqabajhba.jpg"
    	        },
    	"message": "user updated successfully"
    }

## Delete User by UserID

### Request

`DELETE /user/{user_id}`

    localhost:8000/user/id

### Response

    {
    	"data": {
    		"user_id": 1,
    		"nama": "agus",
    		"username": "alta",
    		"email": "alta@alterra.id"
    	        },
    	"message": "delete user by userID successful"
    }

## Create Postingan

### Request

`POST /posts`

    localhost:8000/post

### Response

    {
    	"data": {
    		"posting_id": 1,
    		"pesan": "saya suk12asda7s8li",
    		"foto": "",
    		"user": {
    			"user_id": 1,
    			"nama": "agus",
    			"username": "alta",
    			"foto": "https://res.cloudinary.com/djoxsmzq4/image/upload/v1701989015/sosmed/ldzkcp8txebflvixxvar.jpg"
    		},
    		"comment": null
    	},
    	"message": "success create data"
    }

## Get All Post

### Request

`GET /posts`

    localhost:8000/allpost

### Response

    {
    	"data": [
    		{
    			"posting_id": 1,
    			"pesan": "saya suk12asda7s8li",
    			"foto": "",
    			"user": {
    				"user_id": 1,
    				"nama": "agus",
    				"username": "alta",
    				"foto": "https://res.cloudinary.com/djoxsmzq4/image/upload/v1701991113/sosmed/jdgasy0dgpqcbankmpp9.jpg"
    			},
    			"comment": [
    				{
    					"Comment_id": 1,
    					"pesan": "jika ada yang asdasdasbasdasdaru",
    					"user": {
    						"user_id": 1,
    						"nama": "agus",
    						"username": "alta",
    						"foto": "https://res.cloudinary.com/djoxsmzq4/image/upload/v1701991113/sosmed/jdgasy0dgpqcbankmpp9.jpg"
    					}
    				}
    			]
    		}
    	],
    	"message": "success get all data",
    	"pagination": {
    		"page": 1,
    		"pageSize": 10,
    		"totalPages": 1
    	}
    }

## Update Post by PostID

### Request

`PUT /posts/{post_id}`

    kosong

### Response

    kosong

## Get All Post by UserID

### Request

`GET /posts/{post_id}`

    kosong

### Response

    kosong

## Delete Post by PostID

### Request

`DELETE /posts/{post_id}`

    localhost:8000/delpost

### Response

    {
    	"message": "delete post by userID successful"
    }

## Create Comment

### Request

`POST /comments`

    localhost:8000/comment

### Response

    {
    	"data": {
    		"comment_id": 4,
    		"pesan": "jika ada yang asdasdasbasdasdaru",
    		"User": {
    			"UserID": 1,
    			"Nama": "agus",
    			"UserName": "alta",
    			"Foto": "https://res.cloudinary.com/djoxsmzq4/image/upload/v1701991113/sosmed/jdgasy0dgpqcbankmpp9.jpg"
    		}
    	},
    	"message": "success create data"
    }

## Delete Comment by CommentID

### Request

`DELETE /comments/{comment_id}`

    localhost:8000/comment

### Response

    {
    	"message": "delete comment by commentID successful"
    }
