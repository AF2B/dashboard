CREATE TABLE IF NOT EXISTS ActivityVisitors (
    id SERIAL PRIMARY KEY,
    CurrentValue INTEGER,
    PercentageChange NUMERIC(5,2)
    Date DATE
);

CREATE TABLE IF NOT EXISTS VisitorPerMinute (
    id SERIAL PRIMARY KEY,
    CurrentValue INTEGER,
    PercentageChange NUMERIC(5,2)
    Date DATE
);

CREATE TABLE IF NOT EXISTS ConversionRateLast30Days (
    id SERIAL PRIMARY KEY,
    CurrentValue NUMERIC(5,2),
    PercentageChange NUMERIC(5,2)
    Date DATE
);

CREATE TABLE IF NOT EXISTS BounceRate (
    id SERIAL PRIMARY KEY,
    CurrentValue NUMERIC(5,2),
    PercentageChange NUMERIC(5,2)
    Date DATE
);

CREATE TABLE IF NOT EXISTS BrowserOverview (
    id SERIAL PRIMARY KEY,
    BrowserName VARCHAR(50),
    UsagePercentage NUMERIC(5,2)
    Date DATE
);

CREATE TABLE IF NOT EXISTS UserOverview (
    id SERIAL PRIMARY KEY,
    date DATE,
    ActivityCount INTEGER
    Date DATE
);

CREATE TABLE IF NOT EXISTS MostVisitedPage (
    PageName VARCHAR(255),
    TotalUsers INT,
    BounceRate DECIMAL(5,2)
    Date DATE
);

CREATE TABLE IF NOT EXISTS DevicesOverview (
    DeviceName VARCHAR(255),
    Percentage DECIMAL(5,2),
    Users INT,
    Change VARCHAR(10)
    Date DATE
);

CREATE TABLE IF NOT EXISTS UserLocation (
    CountryName VARCHAR(255),
    Percentage DECIMAL(5,2),
    Users INT
    Date DATE
);

CREATE TABLE IF NOT EXISTS OperatingSystem (
    DeviceType VARCHAR(255),
    OperatingSystem VARCHAR(255),
    Percentage DECIMAL(5,2),
    Users INT
    Date DATE
);

CREATE TABLE IF NOT EXISTS SocialMediaTraffic (
    Network VARCHAR(255),
    Visits INT,
    NewFollowers INT
);
