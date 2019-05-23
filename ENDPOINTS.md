# Endpoints

## Registration

- `/api/register` Register a new user.

```
curl -XPOST -d '{ "name": "heisenberg", "email": "heisenberg@gmail.com", "password": "saymyname!" }' 'http://localhost:8080/api/register'
```

- `/api/login` Login an existing user.

```
curl -XPOST -d '{ "email": "heisenberg@gmail.com", "password": "saymyname!" }' 'http://localhost:8080/api/login'
```

- `/api/logout` Logout current user.

```
curl -XPOST 'http://localhost:8080/api/logout'
```

## User

- `/api/v1/me` GET own user (based on SessionToken).

```
curl -v --cookie "sessionID=710ddd62-4f10-4bf3-8f30-325f7f4a297f" 'http://localhost:8080/api/v1/me'
```


- `/api/v1/users/:id` GET users.

```
curl -v --cookie "sessionID=710ddd62-4f10-4bf3-8f30-325f7f4a297f" 'http://localhost:8080/api/v1/users/1'
```

## Link

- `/api/v1/link/:id` GET link service (based on SessionToken).

```
curl -v --cookie "sessionID=3d266170-9fab-44f6-8694-52a06317b1f6" 'http://localhost:8080/api/v1/link/1'
```


- `/api/v1/link` Create link service (based on SessionToken).

```
curl -v --cookie "sessionID=3d266170-9fab-44f6-8694-52a06317b1f6" 'http://localhost:8080/api/v1/link' -X POST -d 'username=xxx&url=www.xxx.com'
```


- `/api/v1/link/:id` Update link service (based on SessionToken).

```
curl -v --cookie "sessionID=3d266170-9fab-44f6-8694-52a06317b1f6" 'http://localhost:8080/api/v1/link/1' -X PUT -d 'username=xxx&url=www.xxx.com'
```


- `/api/v1/link/:id` Delete link service (based on SessionToken).

```
curl -v --cookie "sessionID=3d266170-9fab-44f6-8694-52a06317b1f6" 'http://localhost:8080/api/v1/link/1' -X DELETE
```

- `/api/v1/links` Get all link service (based on SessionToken).

```
curl -v --cookie "sessionID=9078171b-5b1b-4bcf-bc2e-93a40dd0f47d" 'http://localhost:8080/api/v1/links' -X GET
```
