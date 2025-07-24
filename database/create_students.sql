CREATE TABLE students (
  Timestamp TEXT,                 
  Email TEXT UNIQUE,             
  FullName TEXT,                
  ProgrammeOfStudy TEXT,       
  Faculty TEXT,               
  StudentID TEXT UNIQUE,     
  Level TEXT,                    
  ContactNumber TEXT,            
  EmergencyContact TEXT,
  Presence INTEGER DEFAULT 0,
  OrganiserID TEXT,
  FOREIGN KEY (OrganiserID) REFERENCES organisor(ID)
);
