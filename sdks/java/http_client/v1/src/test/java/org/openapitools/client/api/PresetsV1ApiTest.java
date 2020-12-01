// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.3.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1ListPresetsResponse;
import org.openapitools.client.model.V1Preset;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for PresetsV1Api
 */
@Ignore
public class PresetsV1ApiTest {

    private final PresetsV1Api api = new PresetsV1Api();

    
    /**
     * Create scheduling presets
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createPresetTest() throws ApiException {
        String owner = null;
        V1Preset body = null;
        V1Preset response = api.createPreset(owner, body);

        // TODO: test validations
    }
    
    /**
     * Delete scheduling preset
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deletePresetTest() throws ApiException {
        String owner = null;
        String uuid = null;
        api.deletePreset(owner, uuid);

        // TODO: test validations
    }
    
    /**
     * Get scheduling preset
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getPresetTest() throws ApiException {
        String owner = null;
        String uuid = null;
        V1Preset response = api.getPreset(owner, uuid);

        // TODO: test validations
    }
    
    /**
     * List scheduling presets names
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listPresetNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListPresetsResponse response = api.listPresetNames(owner, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * List scheduling presets
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listPresetsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListPresetsResponse response = api.listPresets(owner, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * Patch scheduling preset
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchPresetTest() throws ApiException {
        String owner = null;
        String presetUuid = null;
        V1Preset body = null;
        V1Preset response = api.patchPreset(owner, presetUuid, body);

        // TODO: test validations
    }
    
    /**
     * Update scheduling preset
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updatePresetTest() throws ApiException {
        String owner = null;
        String presetUuid = null;
        V1Preset body = null;
        V1Preset response = api.updatePreset(owner, presetUuid, body);

        // TODO: test validations
    }
    
}