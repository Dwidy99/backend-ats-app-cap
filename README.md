# ** API CONTRACT GOLANG **

# Login

- Login object

```
 {
   "user_id": int,
   "name": string,
   "username": string,
   "email": string,
   "role": string,
   "token": string,
   "created_at": timestamp
 }
```

## **POST /api/login**

Login to Applicant page and returns the new object.

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **Data Params**

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
        "user_id": int,
        "name": string,
        "username": string,
        "email": string,
        "role": string,
        "token": string,
        "created_at": timestamp
    }
}
```

---

---

# Register

- Register object

```
 {
   "user_id": int,
   "name": string,
   "username": string,
   "email": string,
   "role": string,
   "token": string,
   "created_at": timestamp
 }
```

## **POST /api/applicant/register**

Register to Applicant page and returns the new object.

- **URL Params**  
  None
- **Headers**  
  Content-Type: application/json
- **Data Params for applicant**

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
        "id_user": int
        "name": string,
        "username": string,
        "email": string,
        "role": string,
        "token": string
    }
  }
  ```

---

---

# Users

## **GET /api/employees/users/:id**

Returns the specified user.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
- **Code:** 200
- **Content:**

  ```
  {
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "applicant_id": int,
        "user_id": int,
        "first_name": string,
        "last_name": string,
        "avatar": string,
        "name": string,
        "last_education": string,
        "linkedin_url": string
    }
  }
  ```

- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "User doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

## **PUT /api/employees/users/:id**

Update a User(Applicant) and returns the new object.

- **URL Params**
  None
- **Headers**
  Content-Type: application/json
- **Data Params for employee**

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
        "user_id": int
        "name": string,
        "username": string,
        "email": string,
        "role": string,
        "token": string
    }
  }
  ```

## **DELETE /api/emplooyees/users/:id**

Deletes the specified user.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
  - **Code:** 204
- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "User doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

## **PUT /api/applicants/users/:id**

Update a User(Applicant) and returns the new object.

- **URL Params**
  None
- **Headers**
  Content-Type: application/json
- **Data Params**

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
        "id_applicant": int,
        "user_id": int,
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
    "status": int
}
```

## **GET /api/applicants/experiences**

Returns all Experience in the system.

- **URL Params**
  None
- **Data Params**
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

## **POST /api/applicants/experiences**

Creates a new Experiences for employee and returns the new object.

- **URL Params**
  None
- **Data Params**

```
{
    "company_name": string,
    "role": string,
    "description": text,
    "date_start": date,
    "date_end": date,
    "status": int
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
        "status": int
    }
}
```

## **GET /api/applicants/experiences/:id**

Returns the specified product.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
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
        "status": int
    }
}
```

- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Experience doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

## **DELETE /api/applicants/experiences/:id**

Deletes the specified Experience.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
  - **Code:** 204
- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Experience doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

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

## **GET /api/employees/companies**

Returns all Companies in the system.

- **URL Params**
  None
- **Data Params**
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

## **POST /api/employees/companies**

Creates a new Company for employee and returns the new object.

- **URL Params**
  None
- **Data Params**

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

## **GET /api/employees/companies/:id**

Returns the specified product.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
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

  - **Code:** 404
    **Content:** `{ error : "Company doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

## **PUT /api/employees/companies/:id**

Update a Company and returns the new object.

- **URL Params**
  None
- **Headers**
  Content-Type: application/json
- **Data Params**

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

## **DELETE /api/employees/companies/:id**

Deletes the specified company.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
  - **Code:** 204
- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Company doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

---

---

# Job Categories

- Job Categories object

```

{
    "id": integer,
    "name": string
}

```

## **GET /api/employees/job-categories**

Returns all job-categories in the system.

- **URL Params**
  None
- **Data Params**
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
        {<job_categories_object>},
        {<job_categories_object>},
        {<job_categories_object>}
    }
}
```

## **POST /api/v1/job-categories**

Creates a new Category for employee and returns the new object.

- **URL Params**
  None
- **Data Params**

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
        "job_category_id": integer,
        "name": string
    }
}

```

## **GET /api/employees/job-categories/:id**

Returns the specified product.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
- **Code:** 200
  **Content:**

```
{
    "status": true,
    "message": "ok",
    "errors": null,
    "data": {
        "id": integer,
        "name": string
    }
}
```

- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Job Category doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

## **DELETE /api/employees/job-categories/:id**

Deletes the specified product.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
  - **Code:** 204
- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Job Category doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

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

## **GET /api/employees/jobs**

Returns all jobs list in the system.

- **URL Params**
  None
- **Data Params**
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

## **POST /api/employees/jobs**

Creates a new jobs (for employee and applicant) and returns the new object.

- **URL Params**
  None
- **Data Params**

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

## **GET /api/employees/jobs/:id**

Returns the specified jobs.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
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

- **Error Response:**

  - **Code:** 404
    **Content:** `{ error : "Jobs doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

## **PUT /api/employees/companies/:id**

Update a Company and returns the new object.

- **URL Params**
  None
- **Headers**
  Content-Type: application/json
- **Data Params**

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

## **DELETE /api/employees/jobs/:id**

Deletes the specified jobs.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
  - **Code:** 204
- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Jobs doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

## **GET /api/applicants/jobs**

Returns all jobs available in the system.

- **URL Params**
  None
- **Data Params**
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

## **GET /api/applicants/jobs/:id**

Returns the specified jobs.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
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
        "posted_by": int
    }
}

```

- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Jobs doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

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

## **GET /api/applicant/:id/jobapplication**

Returns all jobs available in the system.

- **URL Params**
  None
- **Data Params**
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
        "jobs": {
            "job_id": integer,
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
            "posted_by": int
        },
        "job_application": {
            "job_applicant_id": int,
            "application_id": int,
            "job_id": int,
            "status": string
        }
    }
}
```

## **POST /api/applicant/jobapplicant**

Applicant apply a job and returns the new object.

- **URL Params**
  None
- **Data Params**

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
        "jobcategory_id": integer,
        "company_id": integer,
        "title": string
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

## **GET /api/applicants/skills**

Returns all skills in the system.

- **URL Params**
  None
- **Data Params**
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

## **POST /api/applicants/skills**

Creates a new Skill for applicant and returns the new object.

- **URL Params**
  None
- **Data Params**

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

## **DELETE /api/applicants/skills/:id**

Deletes the specified skill.

- **URL Params**
  _Required:_ `id=[integer]`
- **Data Params**
  None
- **Headers**
  Content-Type: application/json
  Authorization: Bearer `<OAuth Token>`
- **Success Response:**
  - **Code:** 204
- **Error Response:**
  - **Code:** 404
    **Content:** `{ error : "Skill doesn't exist" }`
    OR
  - **Code:** 401
    **Content:** `{ error : "You are unauthorized to make this request." }`

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
