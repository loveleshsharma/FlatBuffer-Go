# FlatBuffers in Go

```
FlatBuffers is a serialization format from Google. Its very fast in communicating data over network.
Its even faster than JSON and XML.
```

#Getting Started
By following the instructions below, you will get a copy of the project up and running on your local machine for development and testing purposes on FlatBuffers in Go 

### Prequisites for Setup

#### Installation of FlatBuffers (on OSX)
```
$ brew update
$ brew install flatbuffers
``` 
If successful, you will have the flac command accessible from your shell. To verify, execute:
<code>
$ flatc
</code>

#### 1. Writing FlatBuffers schema files
Once you install FlatBuffers, the next step is writing flatbuffers schema definition.
FlatBuffers are defined by schemas. Schemas are plain text files in which we define the data structures we want to serialize.
For Example: <b>empSchema.fbs</b>
```
namespace fbs_schema;

table Employee {
    Eno:int;
    FistName:string;
    LastName:string;
    PhonoNo:string;
    Salary:int;
}

root_type Employee;
```

In the above example, the table represents the Employee object which I want to serialize and deserialize and the root element is Employee. The namespace represents the folder into which the flatbuffer generated files will be placed.

#### 2. Generating Go accessor code from the schema
The next step is to use the <b>flatc</b> compiler to generate Go code for us.
It takes schema file as input and outputs ready-to-use Go code.

In the directory where <b>myempSchema.fbs</b> is located, execute:
```
$ flatc -g empSchema.fbs
```
This will generate Go code in the directory <i>fbs_schema</i>, which was the namespace we mentioned in the schema file. The <b>-g</b> option is used to tell the flatc compiler to generate code in Go.

#### 3. Writing utility functions for the Go accessor code

Once the Go accessor code is generated we need to write the utility functions like serialize and deserialize. Please refer the fbs_utility directory for the code.

#### 4. Ready to Go
After completing the above 3 steps, we are ready to use flatbuffers in our project. We just need to use the utility functions to serialize and deserialize our data. The serialized binary data is very light weight and fast to be sent on network. 
