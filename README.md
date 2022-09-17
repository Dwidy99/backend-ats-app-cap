# ** API CONTRACT GOLANG **

# Login

- Login object

```
 {
   "user_id": integer,
   "name": string,
   "username": string,
   "email": string,
   "role": string,
   "token": string,
   "created_at": timestamp
 }
```

## **POST /login**

Login to Applicant page and returns the new object.

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **Request Body**

```
{
    "email": string,
    "password": string
}
```

- **Success Response:**
- **Code:** 200  
   **Content:**

```
{
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "user_id": integer,
        "name": string,
        "username": string,
        "email": string,
        "role": string,
        "token": string,
        "created_at": timestamp
    }
}
```

- **Success Response:**
- **Code:** 400  
   **Content:**

```
{
    "status": false,
    "message": "Example message",
    "errors": "error",
    "data": null
}
```

---

---

# Register

- Register object

  ```
  {
  "user_id": integer,
  "name": string,
  "username": string,
  "email": string,
  "role": string,
  "token": string,
  "created_at": timestamp
  }
  ```

## **POST /applicants/register**

Register to Applicant page and returns the new object.

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **Request Body for applicant**

  ```
  {
      "name": string,
      "username": string,
      "email": string,
      "password": string
  }
  ```

- **Success Response:**
- **Code:** 200  
   **Content:**

  ```
  {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "id_user": integer,
        "name": string,
        "username": string,
        "email": string,
        "role": string
    }
  }
  ```

- **Success Response:**

- **Code:** 400  
   **Content:**

  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

## **POST /employees/register**

Register to Employees page and returns the new object.(Must Superadmin can be create a new employee)

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **Request Body for applicant**

  ```
  {
      "name": string,
      "username": string,
      "email": string,
      "password": string
  }
  ```

- **Success Response:**
- **Code:** 200  
   **Content:**

  ```
  {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "id_user": integer,
        "name": string,
        "username": string,
        "email": string,
        "role": string
    }
  }
  ```

- **Success Response:**

- **Code:** 400  
   **Content:**

  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

---

---

# Logout

- Logout object

  ```
  {

  }
  ```

## **POST /logout**

Logout and returns no object.

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **No Request Body**

- **Success Response:**
- **Code:** 200  
   **Content:**
  ```
  {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": null
  }
  ```

---

---

# Users

## **GET /employees/users/fetch**

Returns the specified user.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:** 200
- **Content:**

  ```
  {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "user_id": integer,
        "name": string,
        "username": string,
        "email": string,
        "role": string,
        "token": string,
        "contact": string
    }
  }
  ```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **PUT /employees/users/:id**

Update a User(Employee) and returns the new object.

- **URL Params**
  _Required:_ `id=[integer]`
- **Headers**
  Content-Type: application/json
- **Request Body for employee**

  ```
  {
      "name": string,
      "contact": string,
  }
  ```

- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "user_id": integer,
          "name": string,
          "username": string,
          "email": string,
          "role": string,
          "token": string,
          "contact": string
      }
  }
  ```

- **Error Response:**
- **Code:** 400
  **Content:**
  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

## **GET /applicants/users/fetch**

Returns the specified user.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:** 200
- **Content:**

  ```
  {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "user_id": integer,
        "name": string,
        "username": string,
        "email": string,
        "role": string,
        "token": string,
        "contact": string
    }
  }
  ```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **PUT /applicants/users/:id**

Update a User(Applicant) and returns the new object.

- **URL Params**
  _Required:_ `id=[integer]`
- **Headers**
  Content-Type: application/json
- **Request Body**

  ```
  {
      "first_name": string,
      "last_name": string,
      "avatar": string,
      "name": string,
      "last_education": string,
      "linkedin_url": string,
      "github_url": string
  }
  ```

- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
        "id_applicant": integer,
        "user_id": integer,
        "first_name": string,
        "last_name": string,
        "avatar": string,
        "name": string,
        "last_education": string,
        "linkedin_url": string,
        "github_url": string
      }
  }
  ```

  **Error Response:**

- **Code:** 400
  **Content:**
  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

---

---

# Job Experience

- Job Experience object

  ```
  {
      "id_job_experience": integer,
      "applicant_id": integer,
      "company_name": string,
      "role": string,
      "description": text,
      "date_start": date,
      "date_end": date,
      "status": integer
  }
  ```

## **POST /applicants/experiences**

Creates a new Experiences for employee and returns the new object.

- **URL Params**
  None
- **Request Body**

  ```
  {
    "company_name": string,
    "role": string,
    "description": text,
    "date_start": date,
    "date_end": date,
    "status": integer
  }
  ```

- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "job_experience_id": integer,
          "applicant_id": integer,
          "company_name": string,
          "role": string,
          "description": text,
          "date_start": date,
          "date_end": date,
          "status": integer
      }
  }
  ```

  **Error Response:**

  - **Code:** 400
    **Content:**

  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

## **GET /applicants/experiences**

Returns all Experience in the system.

- **URL Params**
  None
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

```
{
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        {<experience_object>},
        {<experience_object>},
        {<experience_object>}
    }
}
```

**Error Response:**

- **Code:** 400
  **Content:**
  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

## **GET /applicants/experiences/:id**

Returns the specified product.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "experience_id": integer,
          "applicant_id": integer,
          "company_name": string,
          "role": string,
          "description": text,
          "date_start": date,
          "date_end": date,
          "status": integer
      }
  }
  ```

- **Error Response:**
- **Code:** 400
  **Content:**
  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

## **DELETE /applicants/experiences/:id**

Deletes the specified Experience.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**

  - **Code:** 200

- **Error Response:**
- **Code:** 400
  **Content:**
  ```
  {
      "status": false,
      "message": "Example message",
      "errors": "error",
      "data": null
  }
  ```

---

---

# Companies

- Company object

```

{
"id": integer,
"name": string,
"email": string,
"address": text,
"contact": string,
"website": string,
"created_at": timestamp
}

```

## **GET /employees/companies**

Returns all Companies in the system.

- **URL Params**
  None
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

```
{
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        {<product_object>},
        {<product_object>},
        {<product_object>}
    }
}
```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **POST /employees/companies**

Creates a new Company for employee and returns the new object.

- **URL Params**
  None
- **Request Body**

  ```
  {
      "name": string,
      "email": string,
      "address": text,
      "contact": string,
      "website": string
  }
  ```

- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**
  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "company_id": integer,
          "name": string,
          "email": string,
          "address": text,
          "contact": string,
          "website": string,
          "created_at": timestamp
      }
  }
  ```

## **GET /employees/companies/:id**

Returns the specified product.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "company_id": integer,
          "name": string,
          "email": string,
          "address": text,
          "contact": string,
          "website": string,
          "created_at": timestamp
      }
  }
  ```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **PUT /employees/companies/:id**

Update a Company and returns the new object.

- **URL Params**
  _Required_ `id=[integer]`
- **Headers**
  Content-Type: application/json
- **Request Body**

  ```
  {
      "name": string,
      "email": string,
      "address": text,
      "contact": string,
      "website": string
  }
  ```

- **Success Response:**
- **Code:** 200
  **Content:**

  ```
    {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "company_id": integer,
        "name": string,
        "email": string,
        "address": text,
        "contact": string,
        "website": string,
        "created_at": timestamp
    }
  }
  ```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **DELETE /employees/companies/:id**

Deletes the specified company.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:**

  ```
  {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": null
  }
  ```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

---

---

# Jobs

- Jobs Object

  ```
  {
      "id": integer,
      "jobcategory_id": integer,
      "company_id": integer,
      "title": string,
      "description": text,
      "location": string,
      "salary": float,
      "type": string,
      "level_of_experience": string,
      "skills": string,
      "date_start": date,
      "date_end": date,
      "created_at": timestamp,
      "posted_by": integer
  }
  ```

## **GET /employees/jobs**

Returns all jobs list in the system.

- **URL Params**
  None
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

```
{
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        {<jobs_object>},
        {<jobs_object>},
        {<jobs_object>}
    }
}
```

## **POST /employees/jobs**

Creates a new jobs and returns the new object.

- **URL Params**
  None
- **Request Body**

  ```
  {
      "company_id": integer,
      "title": string,
      "description": text,
      "location": string,
      "salary": float,
      "type": string,
      "level_of_experience": string,
      "skills": string,
      "date_start": date,
      "date_end": date
  }
  ```

- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "job_id": integer,
          "job_category_id": integer,
          "company_id": integer,
          "title": string,
          "description": text,
          "location": string,
          "salary": float,
          "type": string,
          "level_of_experience": string,
          "skills": string,
          "date_start": date,
          "date_end": date,
          "created_at": timestamp,
          "posted_by": int
      }
  }
  ```

## **GET /employees/jobs/:id**

Returns the specified jobs.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "job_id": integer,
          "job_category_id": integer,
          "company_id": integer,
          "title": string,
          "description": text,
          "location": string,
          "salary": float,
          "type": string,
          "level_of_experience": string,
          "skills": string,
          "date_start": date,
          "date_end": date,
          "created_at": timestamp,
          "posted_by": integer
      }
  }
  ```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **PUT /employees/companies/:id**

Update a Company and returns the new object.

- **URL Params**
  _Required_ `id=[integer]`
- **Headers**
  Content-Type: application/json
- **Request Body**

  ```
  {
    "title": string,
    "description": text,
    "location": string,
    "salary": string,
    "type": string,
    "level_of_experience": string,
    "skills": string,
    "date_start": date,
    "date_end": date,
  }
  ```

- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "company_id": integer,
          "name": string,
          "email": string,
          "address": text,
          "contact": string,
          "website": string,
          "created_at": timestamp
      }
  }
  ```

- **Error Response:**

  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **DELETE /employees/jobs/:id**

Deletes the specified jobs.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:** 200
- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

## **GET /applicants/jobs**

Returns all jobs available in the system.

- **URL Params**
  None
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          {<jobs_object>},
          {<jobs_object>},
          {<jobs_object>}
      }
  }
  ```

## **GET /applicants/jobs/:id**

Returns the specified jobs.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
- **Code:** 200
  **Content:**

```
{
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "job_id": integer,
            "company_id": integer,
            "posted_by": int,
            "title": integer,
            "description": string,
            "location": string,
            "salary": float,
            "type": string,
            "level_of_experience": string,
            "skills": string,
            "company_name": string,
            "address": string,
            "website": string,
            "date_start": date,
            "date_end": date,
            "created_at": timestamp
    }
}
```

- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

---

---

# Job Application

- Job Application Object

  ```
  {
      "id": integer,
      "applicant_id": integer,
      "job_id": integer,
      "status": string,
  }
  ```

## **POST /applicants/jobapplication**

Applicant apply a job and returns the new object.

- **URL Params**
  None
- **Request Body**

  ```
  {
      "status": string
  }
  ```

- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "job_applicant_id": integer,
          "applicant_id": integer,
          "job_id": integer,
          "status": string
      }
  }
  ```

## **POST /employees/jobapplication**

Applicant apply a job and returns the new object.

- **URL Params**
  None
- **Request Body**

  ```
  {
      "status": string
  }
  ```

- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          "job_applicant_id": integer,
          "applicant_id": integer,
          "job_id": integer,
          "status": string
      }
  }
  ```

---

---

# Skills

- Skills object

  ```
  {
      "id": integer,
      "name": string
  }
  ```

## **GET /applicants/skills**

Returns all skills in the system.

- **URL Params**
  None
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          {<skills_object>},
          {<skills_object>},
          {<skills_object>}
      }
  }
  ```

## **POST /applicants/skills**

Creates a new Skill for applicant and returns the new object.

- **URL Params**
  None
- **Request Body**

  ```
  {
      "name": string
  }
  ```

- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
    {
        "status": true,
        "message": "ok",
        "errors": null,
        "data": {
        "skill_id": integer,
        "name": string
        }
    }
  ```

## **DELETE /applicants/skills/:id**

Deletes the specified skill.

- **URL Params**
  _Required:_ `id=[integer]`
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<JWT Token>`
- **Success Response:**
  - **Code:** 200
- **Error Response:**
  - **Code:** 400
    **Content:**
    ```
    {
        "status": false,
        "message": "Example message",
        "errors": "error",
        "data": null
    }
    ```

---

---

# History Applicant

- History Applicant object

  ```
  {
      "applicant_id": integer,
      "user_id": integer,
      "job_applicant_id": integer,
      "job_title": string,
      "job_description": string,
      "job_location": string,
      "job_salary": string,
      "job_type": string,
      "job_level_of_experience": string,
      "status": string
  }
  ```

## **GET /applicants/jobapplicant**

Returns all skills in the system.

- **URL Params**
  None
- **Request Body**
  None
- **Headers**
  Content-Type: application/json
- **Success Response:**
- **Code:** 200
  **Content:**

  ```
  {
      "status": true,
      "message": "ok",
      "errors": null,
      "data": {
          {
            "applicant_id": integer,
            "user_id": integer,
            "job_applicant_id": integer,
            "job_title": string,
            "job_description": string,
            "job_location": string,
            "job_salary": string,
            "job_type": string,
            "job_level_of_experience": string,
            "status": string
          },
          {<skills_object>},
          {<skills_object>}
      }
  }
  ```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```

```
