package netscaler

import (
	"github.com/chiradeep/go-nitro/config/cs"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceNetScalerCsvserver() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCsvserverFunc,
		Read:          readCsvserverFunc,
		Update:        updateCsvserverFunc,
		Delete:        deleteCsvserverFunc,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationhost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authn401": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnvsname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backupvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"casesensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dbprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disableprimaryondown": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmpvsrresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertvserveripport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"l2conn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpolicy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listenpriority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mssqlserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mysqlcharacterset": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlprotocolversion": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlservercapabilities": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mysqlserverversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"precedence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushlabel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushmulticlients": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"redirectportrewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirecturl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtspnat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sobackupaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"somethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sopersistencetimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sothreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stateupdate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	var csvserverName string
	if v, ok := d.GetOk("name"); ok {
		csvserverName = v.(string)
	} else {
		csvserverName = resource.PrefixedUniqueId("tf-csvserver-")
		d.Set("name", csvserverName)
	}
	csvserver := cs.Csvserver{
		Name:                    csvserverName,
		Appflowlog:              d.Get("appflowlog").(string),
		Authentication:          d.Get("authentication").(string),
		Authenticationhost:      d.Get("authenticationhost").(string),
		Authn401:                d.Get("authn401").(string),
		Authnprofile:            d.Get("authnprofile").(string),
		Authnvsname:             d.Get("authnvsname").(string),
		Backupvserver:           d.Get("backupvserver").(string),
		Cacheable:               d.Get("cacheable").(string),
		Casesensitive:           d.Get("casesensitive").(string),
		Clttimeout:              d.Get("clttimeout").(int),
		Comment:                 d.Get("comment").(string),
		Dbprofilename:           d.Get("dbprofilename").(string),
		Disableprimaryondown:    d.Get("disableprimaryondown").(string),
		Downstateflush:          d.Get("downstateflush").(string),
		Httpprofilename:         d.Get("httpprofilename").(string),
		Icmpvsrresponse:         d.Get("icmpvsrresponse").(string),
		Insertvserveripport:     d.Get("insertvserveripport").(string),
		Ipmask:                  d.Get("ipmask").(string),
		Ippattern:               d.Get("ippattern").(string),
		Ipv46:                   d.Get("ipv46").(string),
		L2conn:                  d.Get("l2conn").(string),
		Listenpolicy:            d.Get("listenpolicy").(string),
		Listenpriority:          d.Get("listenpriority").(int),
		Mssqlserverversion:      d.Get("mssqlserverversion").(string),
		Mysqlcharacterset:       d.Get("mysqlcharacterset").(int),
		Mysqlprotocolversion:    d.Get("mysqlprotocolversion").(int),
		Mysqlservercapabilities: d.Get("mysqlservercapabilities").(int),
		Mysqlserverversion:      d.Get("mysqlserverversion").(string),
		Netprofile:              d.Get("netprofile").(string),
		Newname:                 d.Get("newname").(string),
		Port:                    d.Get("port").(int),
		Precedence:              d.Get("precedence").(string),
		Push:                    d.Get("push").(string),
		Pushlabel:               d.Get("pushlabel").(string),
		Pushmulticlients:        d.Get("pushmulticlients").(string),
		Pushvserver:             d.Get("pushvserver").(string),
		Range:                   d.Get("range").(int),
		Redirectportrewrite:     d.Get("redirectportrewrite").(string),
		Redirecturl:             d.Get("redirecturl").(string),
		Rtspnat:                 d.Get("rtspnat").(string),
		Servicetype:             d.Get("servicetype").(string),
		Sobackupaction:          d.Get("sobackupaction").(string),
		Somethod:                d.Get("somethod").(string),
		Sopersistence:           d.Get("sopersistence").(string),
		Sopersistencetimeout:    d.Get("sopersistencetimeout").(int),
		Sothreshold:             d.Get("sothreshold").(int),
		State:                   d.Get("state").(string),
		Stateupdate:             d.Get("stateupdate").(string),
		Tcpprofilename:          d.Get("tcpprofilename").(string),
		Td:                      d.Get("td").(int),
		Vipheader:               d.Get("vipheader").(string),
	}

	_, err := client.AddResource(netscaler.Csvserver.Type(), csvserverName, &csvserver)
	if err != nil {
		return err
	}

	d.SetId(csvserverName)

	err = readCsvserverFunc(d, meta)
	if err != nil {
		log.Printf("?? we just created this csvserver but we can't read it ?? %s", csvserverName)
		return nil
	}
	return nil
}

func readCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Id()
	log.Printf("Reading csvserver state %s", csvserverName)
	data, err := client.FindResource(netscaler.Csvserver.Type(), csvserverName)
	if err != nil {
		log.Printf("Clearing csvserver state %s", csvserverName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("appflowlog", data["appflowlog"])
	d.Set("authentication", data["authentication"])
	d.Set("authenticationhost", data["authenticationhost"])
	d.Set("authn401", data["authn401"])
	d.Set("authnprofile", data["authnprofile"])
	d.Set("authnvsname", data["authnvsname"])
	d.Set("backupvserver", data["backupvserver"])
	d.Set("cacheable", data["cacheable"])
	d.Set("casesensitive", data["casesensitive"])
	d.Set("clttimeout", data["clttimeout"])
	d.Set("comment", data["comment"])
	d.Set("dbprofilename", data["dbprofilename"])
	d.Set("disableprimaryondown", data["disableprimaryondown"])
	d.Set("downstateflush", data["downstateflush"])
	d.Set("httpprofilename", data["httpprofilename"])
	d.Set("icmpvsrresponse", data["icmpvsrresponse"])
	d.Set("insertvserveripport", data["insertvserveripport"])
	d.Set("ipmask", data["ipmask"])
	d.Set("ippattern", data["ippattern"])
	d.Set("ipv46", data["ipv46"])
	d.Set("l2conn", data["l2conn"])
	d.Set("listenpolicy", data["listenpolicy"])
	d.Set("listenpriority", data["listenpriority"])
	d.Set("mssqlserverversion", data["mssqlserverversion"])
	d.Set("mysqlcharacterset", data["mysqlcharacterset"])
	d.Set("mysqlprotocolversion", data["mysqlprotocolversion"])
	d.Set("mysqlservercapabilities", data["mysqlservercapabilities"])
	d.Set("mysqlserverversion", data["mysqlserverversion"])
	d.Set("name", data["name"])
	d.Set("netprofile", data["netprofile"])
	d.Set("newname", data["newname"])
	d.Set("port", data["port"])
	d.Set("precedence", data["precedence"])
	d.Set("push", data["push"])
	d.Set("pushlabel", data["pushlabel"])
	d.Set("pushmulticlients", data["pushmulticlients"])
	d.Set("pushvserver", data["pushvserver"])
	d.Set("range", data["range"])
	d.Set("redirectportrewrite", data["redirectportrewrite"])
	d.Set("redirecturl", data["redirecturl"])
	d.Set("rtspnat", data["rtspnat"])
	d.Set("servicetype", data["servicetype"])
	d.Set("sobackupaction", data["sobackupaction"])
	d.Set("somethod", data["somethod"])
	d.Set("sopersistence", data["sopersistence"])
	d.Set("sopersistencetimeout", data["sopersistencetimeout"])
	d.Set("sothreshold", data["sothreshold"])
	d.Set("state", data["state"])
	d.Set("stateupdate", data["stateupdate"])
	d.Set("tcpprofilename", data["tcpprofilename"])
	d.Set("td", data["td"])
	d.Set("vipheader", data["vipheader"])

	return nil

}

func updateCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] In update func")
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Get("name").(string)

	csvserver := cs.Csvserver{
		Name: d.Get("name").(string),
	}
	if d.HasChange("appflowlog") {
		log.Printf("[DEBUG] Appflowlog has changed for csvserver %s, starting update", csvserverName)
		csvserver.Appflowlog = d.Get("appflowlog").(string)
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG] Authentication has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authentication = d.Get("authentication").(string)
	}
	if d.HasChange("authenticationhost") {
		log.Printf("[DEBUG] Authenticationhost has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authenticationhost = d.Get("authenticationhost").(string)
	}
	if d.HasChange("authn401") {
		log.Printf("[DEBUG] Authn401 has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authn401 = d.Get("authn401").(string)
	}
	if d.HasChange("authnprofile") {
		log.Printf("[DEBUG] Authnprofile has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authnprofile = d.Get("authnprofile").(string)
	}
	if d.HasChange("authnvsname") {
		log.Printf("[DEBUG] Authnvsname has changed for csvserver %s, starting update", csvserverName)
		csvserver.Authnvsname = d.Get("authnvsname").(string)
	}
	if d.HasChange("backupvserver") {
		log.Printf("[DEBUG] Backupvserver has changed for csvserver %s, starting update", csvserverName)
		csvserver.Backupvserver = d.Get("backupvserver").(string)
	}
	if d.HasChange("cacheable") {
		log.Printf("[DEBUG] Cacheable has changed for csvserver %s, starting update", csvserverName)
		csvserver.Cacheable = d.Get("cacheable").(string)
	}
	if d.HasChange("casesensitive") {
		log.Printf("[DEBUG] Casesensitive has changed for csvserver %s, starting update", csvserverName)
		csvserver.Casesensitive = d.Get("casesensitive").(string)
	}
	if d.HasChange("clttimeout") {
		log.Printf("[DEBUG] Clttimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Clttimeout = d.Get("clttimeout").(int)
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG] Comment has changed for csvserver %s, starting update", csvserverName)
		csvserver.Comment = d.Get("comment").(string)
	}
	if d.HasChange("dbprofilename") {
		log.Printf("[DEBUG] Dbprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Dbprofilename = d.Get("dbprofilename").(string)
	}
	if d.HasChange("disableprimaryondown") {
		log.Printf("[DEBUG] Disableprimaryondown has changed for csvserver %s, starting update", csvserverName)
		csvserver.Disableprimaryondown = d.Get("disableprimaryondown").(string)
	}
	if d.HasChange("downstateflush") {
		log.Printf("[DEBUG] Downstateflush has changed for csvserver %s, starting update", csvserverName)
		csvserver.Downstateflush = d.Get("downstateflush").(string)
	}
	if d.HasChange("httpprofilename") {
		log.Printf("[DEBUG] Httpprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Httpprofilename = d.Get("httpprofilename").(string)
	}
	if d.HasChange("icmpvsrresponse") {
		log.Printf("[DEBUG] Icmpvsrresponse has changed for csvserver %s, starting update", csvserverName)
		csvserver.Icmpvsrresponse = d.Get("icmpvsrresponse").(string)
	}
	if d.HasChange("insertvserveripport") {
		log.Printf("[DEBUG] Insertvserveripport has changed for csvserver %s, starting update", csvserverName)
		csvserver.Insertvserveripport = d.Get("insertvserveripport").(string)
	}
	if d.HasChange("ipmask") {
		log.Printf("[DEBUG] Ipmask has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipmask = d.Get("ipmask").(string)
	}
	if d.HasChange("ippattern") {
		log.Printf("[DEBUG] Ippattern has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ippattern = d.Get("ippattern").(string)
	}
	if d.HasChange("ipv46") {
		log.Printf("[DEBUG] Ipv46 has changed for csvserver %s, starting update", csvserverName)
		csvserver.Ipv46 = d.Get("ipv46").(string)
	}
	if d.HasChange("l2conn") {
		log.Printf("[DEBUG] L2conn has changed for csvserver %s, starting update", csvserverName)
		csvserver.L2conn = d.Get("l2conn").(string)
	}
	if d.HasChange("listenpolicy") {
		log.Printf("[DEBUG] Listenpolicy has changed for csvserver %s, starting update", csvserverName)
		csvserver.Listenpolicy = d.Get("listenpolicy").(string)
	}
	if d.HasChange("listenpriority") {
		log.Printf("[DEBUG] Listenpriority has changed for csvserver %s, starting update", csvserverName)
		csvserver.Listenpriority = d.Get("listenpriority").(int)
	}
	if d.HasChange("mssqlserverversion") {
		log.Printf("[DEBUG] Mssqlserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mssqlserverversion = d.Get("mssqlserverversion").(string)
	}
	if d.HasChange("mysqlcharacterset") {
		log.Printf("[DEBUG] Mysqlcharacterset has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlcharacterset = d.Get("mysqlcharacterset").(int)
	}
	if d.HasChange("mysqlprotocolversion") {
		log.Printf("[DEBUG] Mysqlprotocolversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlprotocolversion = d.Get("mysqlprotocolversion").(int)
	}
	if d.HasChange("mysqlservercapabilities") {
		log.Printf("[DEBUG] Mysqlservercapabilities has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlservercapabilities = d.Get("mysqlservercapabilities").(int)
	}
	if d.HasChange("mysqlserverversion") {
		log.Printf("[DEBUG] Mysqlserverversion has changed for csvserver %s, starting update", csvserverName)
		csvserver.Mysqlserverversion = d.Get("mysqlserverversion").(string)
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG] Name has changed for csvserver %s, starting update", csvserverName)
		csvserver.Name = d.Get("name").(string)
	}
	if d.HasChange("netprofile") {
		log.Printf("[DEBUG] Netprofile has changed for csvserver %s, starting update", csvserverName)
		csvserver.Netprofile = d.Get("netprofile").(string)
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG] Newname has changed for csvserver %s, starting update", csvserverName)
		csvserver.Newname = d.Get("newname").(string)
	}
	if d.HasChange("port") {
		log.Printf("[DEBUG] Port has changed for csvserver %s, starting update", csvserverName)
		csvserver.Port = d.Get("port").(int)
	}
	if d.HasChange("precedence") {
		log.Printf("[DEBUG] Precedence has changed for csvserver %s, starting update", csvserverName)
		csvserver.Precedence = d.Get("precedence").(string)
	}
	if d.HasChange("push") {
		log.Printf("[DEBUG] Push has changed for csvserver %s, starting update", csvserverName)
		csvserver.Push = d.Get("push").(string)
	}
	if d.HasChange("pushlabel") {
		log.Printf("[DEBUG] Pushlabel has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushlabel = d.Get("pushlabel").(string)
	}
	if d.HasChange("pushmulticlients") {
		log.Printf("[DEBUG] Pushmulticlients has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushmulticlients = d.Get("pushmulticlients").(string)
	}
	if d.HasChange("pushvserver") {
		log.Printf("[DEBUG] Pushvserver has changed for csvserver %s, starting update", csvserverName)
		csvserver.Pushvserver = d.Get("pushvserver").(string)
	}
	if d.HasChange("range") {
		log.Printf("[DEBUG] Range has changed for csvserver %s, starting update", csvserverName)
		csvserver.Range = d.Get("range").(int)
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG] Redirectportrewrite has changed for csvserver %s, starting update", csvserverName)
		csvserver.Redirectportrewrite = d.Get("redirectportrewrite").(string)
	}
	if d.HasChange("redirecturl") {
		log.Printf("[DEBUG] Redirecturl has changed for csvserver %s, starting update", csvserverName)
		csvserver.Redirecturl = d.Get("redirecturl").(string)
	}
	if d.HasChange("rtspnat") {
		log.Printf("[DEBUG] Rtspnat has changed for csvserver %s, starting update", csvserverName)
		csvserver.Rtspnat = d.Get("rtspnat").(string)
	}
	if d.HasChange("servicetype") {
		log.Printf("[DEBUG] Servicetype has changed for csvserver %s, starting update", csvserverName)
		csvserver.Servicetype = d.Get("servicetype").(string)
	}
	if d.HasChange("sobackupaction") {
		log.Printf("[DEBUG] Sobackupaction has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sobackupaction = d.Get("sobackupaction").(string)
	}
	if d.HasChange("somethod") {
		log.Printf("[DEBUG] Somethod has changed for csvserver %s, starting update", csvserverName)
		csvserver.Somethod = d.Get("somethod").(string)
	}
	if d.HasChange("sopersistence") {
		log.Printf("[DEBUG] Sopersistence has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sopersistence = d.Get("sopersistence").(string)
	}
	if d.HasChange("sopersistencetimeout") {
		log.Printf("[DEBUG] Sopersistencetimeout has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sopersistencetimeout = d.Get("sopersistencetimeout").(int)
	}
	if d.HasChange("sothreshold") {
		log.Printf("[DEBUG] Sothreshold has changed for csvserver %s, starting update", csvserverName)
		csvserver.Sothreshold = d.Get("sothreshold").(int)
	}
	if d.HasChange("state") {
		log.Printf("[DEBUG] State has changed for csvserver %s, starting update", csvserverName)
		csvserver.State = d.Get("state").(string)
	}
	if d.HasChange("stateupdate") {
		log.Printf("[DEBUG] Stateupdate has changed for csvserver %s, starting update", csvserverName)
		csvserver.Stateupdate = d.Get("stateupdate").(string)
	}
	if d.HasChange("tcpprofilename") {
		log.Printf("[DEBUG] Tcpprofilename has changed for csvserver %s, starting update", csvserverName)
		csvserver.Tcpprofilename = d.Get("tcpprofilename").(string)
	}
	if d.HasChange("td") {
		log.Printf("[DEBUG] Td has changed for csvserver %s, starting update", csvserverName)
		csvserver.Td = d.Get("td").(int)
	}
	if d.HasChange("vipheader") {
		log.Printf("[DEBUG] Vipheader has changed for csvserver %s, starting update", csvserverName)
		csvserver.Vipheader = d.Get("vipheader").(string)
	}

	_, err := client.UpdateResource(netscaler.Csvserver.Type(), csvserverName, &csvserver)
	if err != nil {
		return fmt.Errorf("Error updating csvserver %s", csvserverName)
	}
	return readCsvserverFunc(d, meta)
}

func deleteCsvserverFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	csvserverName := d.Id()
	err := client.DeleteResource(netscaler.Csvserver.Type(), csvserverName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
