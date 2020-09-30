

### CLIENT REQUESTS (CURL)

##### THE CLIENT REQUEST WITH "0" as Test value

	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/helloapi/fibonacci?num=0 


##### THE RESPONSE FROM SERVER
	
	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 30 Sep 2020 02:17:30 GMT
	Content-Length: 172

	{
	    "id": "8edaf307-22a0-2f5d-0768-5de138f8148e",
	    "inputNum": 0,
	    "fiboNum": 0,
	    "created_at": "2020-09-30T04:17:30.252482907+02:00",
	    "updated_at": "2020-09-30T04:17:30.252482974+02:00"
	}

##### THE CLIENT REQUEST WITH "1" as Test value

	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/helloapi/fibonacci?num=1


##### THE RESPONSE FROM SERVER

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 30 Sep 2020 02:21:00 GMT
	Content-Length: 171

	{
	    "id": "41706a02-2924-218e-f8f5-2c94ba82c554",
	    "inputNum": 1,
	    "fiboNum": 1,
	    "created_at": "2020-09-30T04:21:00.98951589+02:00",
	    "updated_at": "2020-09-30T04:21:00.989515961+02:00"
	}


##### THE CLIENT REQUEST WITH "2" as Test value

	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/helloapi/fibonacci?num=2


##### THE RESPONSE FROM SERVER

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 30 Sep 2020 02:22:01 GMT
	Content-Length: 172

	{
	    "id": "49e99f38-4833-b52e-4573-f91fcd5ad549",
	    "inputNum": 2,
	    "fiboNum": 1,
	    "created_at": "2020-09-30T04:22:01.278997923+02:00",
	    "updated_at": "2020-09-30T04:22:01.279000634+02:00"
	}

##### THE CLIENT REQUEST WITH "10" as Test value

	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/helloapi/fibonacci?num=10

##### THE RESPONSE FROM SERVER

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 30 Sep 2020 02:23:22 GMT
	Content-Length: 174

	{
	    "id": "acfa2715-c3c8-08d0-f94e-784bcc4eced8",
	    "inputNum": 10,
	    "fiboNum": 55,
	    "created_at": "2020-09-30T04:23:22.276220429+02:00",
	    "updated_at": "2020-09-30T04:23:22.276220496+02:00"
	}
	
##	Planning Poker

	
##### THE CLIENT REQUEST WITH "11" as Test value

	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/helloapi/fibonacci?num=11
	

##### THE RESPONSE FROM SERVER	
	
	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 30 Sep 2020 02:33:09 GMT
	Content-Length: 174

	{
	    "id": "7b825d0b-2a71-a8b6-dfa4-efae64254bc6",
	    "inputNum": 11,
	    "fiboNum": 89,
	    "created_at": "2020-09-30T04:33:09.995112422+02:00",
	    "updated_at": "2020-09-30T04:33:09.995112603+02:00"
	}
	

##### THE CLIENT REQUEST WITH "12" as Test value

	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/helloapi/fibonacci?num=12


##### THE RESPONSE FROM SERVER	

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 30 Sep 2020 02:35:17 GMT
	Content-Length: 175

	{
	    "id": "a256cff4-f9a0-92dc-bbfe-ecd8c742456e",
	    "inputNum": 12,
	    "fiboNum": 100,
	    "created_at": "2020-09-30T04:35:17.370313984+02:00",
	    "updated_at": "2020-09-30T04:35:17.370314087+02:00"
	}
	
##### THE CLIENT REQUEST WITH "12" as Test value

	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/helloapi/fibonacci?num=45312


##### THE RESPONSE FROM SERVER

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 30 Sep 2020 02:37:32 GMT
	Content-Length: 178

	{
	    "id": "fb879981-e848-c925-9435-c56060b25b58",
	    "inputNum": 45312,
	    "fiboNum": 100,
	    "created_at": "2020-09-30T04:37:32.958233036+02:00",
	    "updated_at": "2020-09-30T04:37:32.958233124+02:00"
	}
