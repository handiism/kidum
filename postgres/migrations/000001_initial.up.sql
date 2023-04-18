CREATE TYPE "FormType" AS enum ('INDIVIDUAL', 'GROUP');

CREATE TYPE "Gender" AS enum ('MALE', 'FEMALE');

CREATE TYPE "Role" AS enum ('VERIFIER', 'MONITOR');

CREATE TYPE "CredentialType" AS enum ('NATIONAL_ID', 'DRIVER_LICENSE');

CREATE TYPE "FormStatus" AS enum (
    'UNPROCCESSED',
    'DATA_INVALID',
    'DATA_VALID',
    'DONE'
);

CREATE TABLE IF NOT EXISTS "Form" (
    "id" SERIAL PRIMARY KEY,
    "code" VARCHAR(8) UNIQUE,
    "agreement" BOOLEAN DEFAULT FALSE NOT NULL,
    "ticket" TEXT,
    "description" TEXT,
    -- 
    "status" "FormStatus" DEFAULT 'UNPROCCESSED' :: "FormStatus" NOT NULL,
    "type" "FormType" NOT NULL,
    -- 
    "createdAt" TIMESTAMPTZ(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMPTZ(3),
    "deleted_at" TIMESTAMPTZ(3),
    -- 
    "destinationRouteId" INTEGER,
    "adminScannerId" INTEGER,
    "adminValidatorId" INTEGER,
    "contactId" INTEGER
);

CREATE TABLE IF NOT EXISTS "Admin" (
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR(63) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "role" "Role" DEFAULT 'VERIFIER' :: "Role" NOT NULL
);

CREATE TABLE IF NOT EXISTS "DestinationRoute" (
    "id" SERIAL PRIMARY KEY,
    "order" INTEGER NOT NULL,
    -- 
    "routeId" INTEGER,
    "destinationId" INTEGER
);

CREATE TABLE IF NOT EXISTS "Route" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "quota" INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS "Destination" (
    "id" SERIAL PRIMARY KEY,
    "city" VARCHAR(255) NOT NULL,
    "province" VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "Contact" (
    "id" SERIAL PRIMARY KEY,
    "number" VARCHAR(63),
    "emergencyNumber" VARCHAR(63),
    "email" VARCHAR(63)
);

CREATE TABLE IF NOT EXISTS "User" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "age" INTEGER NOT NULL,
    "birthDate" DATE NOT NULL,
    "address" TEXT NOT NULL,
    -- 
    "gender" "Gender" DEFAULT 'MALE' :: "Gender" NOT NULL,
    -- 
    "deleted_at" TIMESTAMPTZ(3),
    "updatedAt" TIMESTAMPTZ(3),
    "createdAt" TIMESTAMPTZ(3) DEFAULT CURRENT_TIMESTAMP NOT NULL,
    -- 
    "formId" INTEGER,
    "credentialId" INTEGER
);

CREATE TABLE IF NOT EXISTS "Credential" (
    "id" SERIAL PRIMARY KEY,
    "image" TEXT NOT NULL,
    "number" VARCHAR(63) NOT NULL UNIQUE,
    -- 
    "type" "CredentialType" NOT NULL
);