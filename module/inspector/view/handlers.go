package inspectorview

import (
	"fmt"
	inspectorv1 "sdm/pb/sdm/inspector/v1"

	"github.com/gofiber/fiber/v2"
)

func (s *InspectorRouter) DatabaseList(c *fiber.Ctx) error {
	inspector := c.Locals("inspector").(inspectorv1.InspectorServiceServer)
	ctx := c.UserContext()

	res, err := inspector.ListDatabases(ctx, &inspectorv1.ListDatabasesRequest{})
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (s *InspectorRouter) SchemaList(c *fiber.Ctx) error {
	inspector := c.Locals("inspector").(inspectorv1.InspectorServiceServer)
	ctx := c.UserContext()

	res, err := inspector.ListSchemas(ctx, &inspectorv1.ListSchemasRequest{})
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (s *InspectorRouter) ObjectList(c *fiber.Ctx) error {
	inspector := c.Locals("inspector").(inspectorv1.InspectorServiceServer)
	ctx := c.UserContext()

	res, err := inspector.ListObjects(ctx, &inspectorv1.ListObjectsRequest{})
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (s *InspectorRouter) SchemaTableList(c *fiber.Ctx) error {
	inspector := c.Locals("inspector").(inspectorv1.InspectorServiceServer)
	ctx := c.UserContext()

	schemaid := c.Params("schemaid")

	res, err := inspector.ListTables(ctx, &inspectorv1.ListTablesRequest{
		Schema: schemaid,
	})
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (s *InspectorRouter) TableInfo(c *fiber.Ctx) error {
	inspector := c.Locals("inspector").(inspectorv1.InspectorServiceServer)
	ctx := c.UserContext()

	schemaid := c.Params("schemaid")
	tableid := c.Params("tableid")

	res, err := inspector.GetTableInfo(ctx, &inspectorv1.GetTableInfoRequest{
		Schema: schemaid,
		Table:  tableid,
	})
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (s *InspectorRouter) Query(c *fiber.Ctx) error {
	inspector := c.Locals("inspector").(inspectorv1.InspectorServiceServer)
	ctx := c.UserContext()

	req := &inspectorv1.QueryRequest{}
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).SendString(fmt.Sprintf("error parsing %v", err))
	}

	res, err := inspector.Query(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
