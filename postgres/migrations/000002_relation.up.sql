ALTER TABLE IF EXISTS "Form"
ADD CONSTRAINT "FkFormDestinationRoute"
FOREIGN KEY ("destinationRouteId")
REFERENCES "DestinationRoute"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;

ALTER TABLE IF EXISTS "Form"
ADD CONSTRAINT "FkFormAdminScanner"
FOREIGN KEY ("adminScannerId")
REFERENCES "Admin"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;

ALTER TABLE IF EXISTS "Form"
ADD CONSTRAINT "FkFormAdminValidator"
FOREIGN KEY ("adminValidatorId")
REFERENCES "Admin"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;

ALTER TABLE IF EXISTS "Form"
ADD CONSTRAINT "FkFormContact"
FOREIGN KEY ("contactId")
REFERENCES "Contact"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;

ALTER TABLE IF EXISTS "DestinationRoute"
ADD CONSTRAINT "FkRoute"
FOREIGN KEY ("routeId")
REFERENCES "Route"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;

ALTER TABLE IF EXISTS "DestinationRoute"
ADD CONSTRAINT "FkDestination"
FOREIGN KEY ("destinationId")
REFERENCES "Destination"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;


ALTER TABLE IF EXISTS "User"
ADD CONSTRAINT "FkUserForm"
FOREIGN KEY ("formId")
REFERENCES "Form"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;

ALTER TABLE IF EXISTS "User"
ADD CONSTRAINT "FkUserCredential"
FOREIGN KEY ("credentialId")
REFERENCES "Credential"("id")
ON DELETE SET NULL
ON UPDATE CASCADE;

