# RESTAPI schema for list-of-expenses

## Open Endpoints

open endpoints no require Authentication.
* Create spent `POST /api/spent`

    **Data constraints**
        
        ```json
        {
            "name": "names of expenses here in string",
            "amount": "amount here in integer"
        }
        ```
    **Code** : `201 Created`
    
    **Content example** : Return ID created spent
    
    
* Show spent `GET /api/spent/{:id}`
  
  **Success Responses Code** : `200 OK`
  
  **Content example** : show spent by his id
  
  
* Update spent `PUT /api/spent/{:id}`

    **Data constraints**
            
            ```json
            {
                "name": "names of expenses here in string",
                "amount": "amount here in integer"
            }
            ```
    **Success Responses Code** : `200 OK`
      
    **Content example** : Return name and amount
* Delete spent `DELETE /api/spent/{:id}`

* Show all spent by date start/end `GET /api/spents`

    **Data constraints**
    
    ```query
    date_from=%s&date_to=%s
    ```


