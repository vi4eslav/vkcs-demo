package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
	"time"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider})
}

const DefaultValue = "a"

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"my_option": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MYAPP_OPTION1", DefaultValue),
				Description: "The myapp optiom 1.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"myapp_hello": resourceHello(),
		},
	}
	provider.ConfigureContextFunc = func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return &HelloApp{myOption: d.Get("my_option").(string)}, nil // Used for all CRUD providers
	}

	return provider
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func resourceHello() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceHelloCreate,
		ReadContext:   resourceHelloRead,
		UpdateContext: resourceHelloUpdate,
		DeleteContext: resourceHelloDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"content_from_conf": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The content from config.",
			},
			"content_from_app": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "The content from app.",
			},
		},
	}
}

func resourceHelloCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*HelloApp)
	var id, content string
	var errStatus error
	// *******
	_ = resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		id, content, errStatus = client.Create(ctx)
		if errStatus != nil {
			log.Printf("[DEBUG] Retrying after error: %s", errStatus)
			return resource.RetryableError(errStatus)
			//return resource.NonRetryableError(errStatus)
		}
		return nil
	})
	// *******
	d.SetId(id)
	d.Set("content_from_app", content)
	return nil
}

func resourceHelloRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*HelloApp)
	content, err := client.Get(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("content_from_app", content)
	return nil
}

func resourceHelloUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*HelloApp)
	content, err := client.Update(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("content_from_app", content)

	return nil
}

func resourceHelloDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*HelloApp)
	err := client.Delete(ctx, d.Id())
	return diag.FromErr(err)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type HelloApp struct {
	myOption string
}

func (ha *HelloApp) Create(ctx context.Context) (string, string, error) {
	u, err := uuid.GenerateUUID()
	return u, fmt.Sprintf("Hello: %s!", u), err
}

func (ha *HelloApp) Update(ctx context.Context, id string) (string, error) {
	return fmt.Sprintf("Hello: %s!", id), nil
}

func (ha *HelloApp) Get(ctx context.Context, id string) (string, error) {
	// return fmt.Sprintf("Hello: %s!", id), err
	u, err := uuid.GenerateUUID()
	return fmt.Sprintf("Hello: %s!", u), err
}

func (ha *HelloApp) Delete(ctx context.Context, id string) error {
	return nil
}
