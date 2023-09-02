DROP TABLE Users 
(
    Id INT NOT NULL AUTO_INCREMENT,
    Name VARCHAR(255),
    CONSTRAINT PK_User_Id PRIMARY KEY (Id)
);

DROP TABLE Dependencies 
(
    Id INT NOT NULL AUTO_INCREMENT,
    UserId INT NOT NULL,
    Segment VARCHAR(512),
    CONSTRAINT PK_Segment_Id PRIMARY KEY (Id),
    CONSTRAINT FK_Segment_User FOREIGN KEY (UserId) REFERENCES Users (Id) ON DELETE CASCADE ON UPDATE CASCADE
);

DROP TABLE Segments 
(
    Id INT NOT NULL AUTO_INCREMENT,
    Segment VARCHAR(255),
    CONSTRAINT PK_Segment_Id PRIMARY KEY (Id)
);