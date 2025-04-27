package provider

import (
	"context"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/es6kr/terraform-provider-mac/internal/launchctl"
	"github.com/es6kr/terraform-provider-mac/internal/xerrors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const serviceIDPrefix = "service:"

func serviceID(label string) string {
	return serviceIDPrefix + label
}

func labelFromServiceID(id string) string {
	return strings.TrimPrefix(id, serviceIDPrefix)
}

func resourceService() *schema.Resource {
	return &schema.Resource{
		Description: "`mac_service` manages a service using [launchctl](https://www.manpagez.com/man/1/launchctl/).\n\n" +
			"It controls macOS services by loading, unloading, and optionally starting them via `launchctl`.",
		CreateContext: resourceServiceCreate,
		ReadContext:   resourceServiceRead,
		DeleteContext: resourceServiceDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Internal ID of the service resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"label": {
				Description: "The label of the launchctl service, usually defined in the plist.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"plist": {
				Description: "Absolute path to the service plist file.",
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
			},
			"start": {
				Description: "Whether to immediately start the service after loading.",
				Type:        schema.TypeBool,
				Default:     false,
				ForceNew:    true,
				Optional:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceServiceCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	label := data.Get("label").(string)
	plist := data.Get("plist").(string)
	start := data.Get("start").(bool)

	if err := launchctl.Bootstrap(ctx, plist); err != nil {
		return xerrors.ToDiags(err)
	}

	if start {
		if err := launchctl.Kickstart(ctx, label); err != nil {
			return xerrors.ToDiags(err)
		}
	}

	data.SetId(serviceID(label))

	return resourceServiceRead(ctx, data, meta)
}

func resourceServiceRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	label := labelFromServiceID(data.Id())

	found, err := launchctl.IsLoaded(ctx, label)
	if err != nil {
		return xerrors.ToDiags(err)
	}

	if !found {
		data.SetId("")
	}

	return nil
}

func resourceServiceDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	label := labelFromServiceID(data.Id())

	if err := launchctl.Bootout(ctx, label); err != nil {
		if errors.Is(err, xerrors.ErrNotInstalled) {
			// Already removed
			data.SetId("")
			return nil
		}
		return xerrors.ToDiags(err)
	}

	data.SetId("")

	return nil
}
