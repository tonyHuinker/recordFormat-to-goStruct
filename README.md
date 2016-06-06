# recordFormat-to-goStruct

utility to convert ExtraHop Record Formats (json) into Golang structs when using the ExtraHop REST api to get EXA data  

Instructions.

1. Copy ExtraHop Record Format into file
2. Run program
3. Program will ask location of Format file
4. Enter the name you would like to use for Golang struct

*Note.  There are times when the same 'name' can be used twice, which is not allowed in a Go struct.  If this happens and they are the same 'type' (like both are strings, or both are devices), then just delete one of them as they cannot both be returned simultaneously.  If however they have different types.. please contact extrahop@support.com and let us know!*

*Example Input  and Output Files provided.  Example is ExtraHop's built in ica_close record*
