CREATE TABLE emails (
  ID INTEGER PRIMARY KEY AUTOINCREMENT,  
  Email TEXT UNIQUE NOT NULL,            
  Password TEXT NOT NULL,
  AppPassword TEXT NOT NULL,
  Sent INTEGER DEFAULT 0 
);
