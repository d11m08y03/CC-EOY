CREATE TABLE organisors (
  ID INTEGER PRIMARY KEY AUTOINCREMENT,  
  Name TEXT NOT NULL,                    
  Email TEXT UNIQUE NOT NULL,            
  Password TEXT NOT NULL                 
);
