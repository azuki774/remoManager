#!/bin/bash
curl -X GET "https://api.nature.global/1/devices" -H "accept: application/json" -k --header "Authorization: Bearer ${TOKEN}"
