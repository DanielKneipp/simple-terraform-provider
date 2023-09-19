package provider

import (
	client "terraform-provider-python-app/api"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceList() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"numbers": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of numbers in the server",
				ForceNew:    false,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
		Create: resourceCreateList,
		Read:   resourceReadList,
		Update: resourceUpdateList,
		Delete: resourceDeleteList,
		Exists: resourceExistsList,
	}
}

func resourceCreateList(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	numbers_interface := d.Get("numbers").([]interface{})

	numbers := make([]int, len(numbers_interface))

	for i, num_interface := range numbers_interface {
		numbers[i] = num_interface.(int)
	}

	err := apiClient.RemoveAll()
	if err != nil {
		return err
	}

	for _, num := range numbers {
		apiClient.AddNumber(num)
	}

	d.SetId("id")
	return nil
}

func resourceReadList(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	numbers, err := apiClient.GetAll()
	if err != nil {
		return err
	}

	d.SetId("id")
	d.Set("numbers", numbers)

	return nil
}

func resourceUpdateList(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	numbers_interface := d.Get("numbers").([]interface{})

	numbers := make([]int, len(numbers_interface))

	for i, num_interface := range numbers_interface {
		numbers[i] = num_interface.(int)
	}

	err := apiClient.RemoveAll()
	if err != nil {
		return err
	}

	for _, num := range numbers {
		apiClient.AddNumber(num)
	}

	return nil
}

func resourceDeleteList(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	err := apiClient.RemoveAll()
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceExistsList(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	numbers, err := apiClient.GetAll()
	if err != nil {
		return false, err
	}

	if len(numbers) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
