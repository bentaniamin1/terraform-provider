package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	tfpath "github.com/hashicorp/terraform-plugin-framework/path"
	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *UserResource) Create(ctx context.Context, req tfresource.CreateRequest, resp *tfresource.CreateResponse) {
	var state UserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := r.client.Post(r.endpoint+state.Id.ValueString(), "application/text", bytes.NewBuffer([]byte(state.Name.ValueString())))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error sending request: %s", err))
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("HTTP Error", fmt.Sprintf("Received non-OK HTTP status: %s", response.Status))
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Response Body", fmt.Sprintf("Could not read response body: %s", err))
		return
	}
	state.Name = tftypes.StringValue(string(body))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserResource) Read(ctx context.Context, req tfresource.ReadRequest, resp *tfresource.ReadResponse) {
	var state UserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	response, err := r.client.Get(r.endpoint + state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read user, got error: %s", err))
		return
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		resp.State.RemoveResource(ctx)
		return
	}
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			resp.Diagnostics.AddError("Error reading response body", err.Error())
			return
		}
		state.Name = tftypes.StringValue(string(bodyBytes))
		resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
		return
	}
	resp.Diagnostics.AddError("HTTP Error", fmt.Sprintf("Received bad HTTP status: %s", response.Status))
}

func (r *UserResource) Delete(ctx context.Context, req tfresource.DeleteRequest, resp *tfresource.DeleteResponse) {
	var data UserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request, err := http.NewRequest(http.MethodDelete, r.endpoint+data.Id.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Request Creation Failed", fmt.Sprintf("Could not create HTTP request: %s", err))
		return
	}
	response, err := r.client.Do(request)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete user, got error: %s", err))
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("HTTP Error", fmt.Sprintf("Received non-OK HTTP status: %s", response.Status))
		return
	}
	data.Id = tftypes.StringValue("")
	data.Name = tftypes.StringValue("")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserResource) Update(ctx context.Context, req tfresource.UpdateRequest, resp *tfresource.UpdateResponse) {
	var state UserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	webserviceCall, err := http.NewRequest("PUT", r.endpoint+state.Id.ValueString(), bytes.NewBuffer([]byte(state.Name.ValueString())))
	if err != nil {
		resp.Diagnostics.AddError("Go Error", fmt.Sprintf("Error sending request: %s", err))
	}
	response, err := r.client.Do(webserviceCall)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error sending request: %s", err))
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		resp.Diagnostics.AddError("HTTP Error", fmt.Sprintf("Received non-OK HTTP status: %s", response.Status))
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Read Response Body", fmt.Sprintf("Could not read response body: %s", err))
		return
	}
	state.Name = tftypes.StringValue(string(body))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserResource) ImportState(ctx context.Context, req tfresource.ImportStateRequest, resp *tfresource.ImportStateResponse) {
	tfresource.ImportStatePassthroughID(ctx, tfpath.Root("id"), req, resp)
}
