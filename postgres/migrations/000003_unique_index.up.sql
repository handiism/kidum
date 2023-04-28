ALTER TABLE IF EXISTS "DestinationRoute"
ADD CONSTRAINT "UqRouteDestinationOrder"
UNIQUE ("routeId","destinationId","order");