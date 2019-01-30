// Tensorflow Serving Go client for the inception model
// First of all compile the proto files:
// git clone --recursive https://github.com/tensorflow/serving.git
package main
import (
    "context"
    "log"
    pb "github.com/jnummelin/go-inception-client/tensorflow_serving/apis"
    tf_core_framework "github.com/bmwhopper/go-inception-client/tensorflow/core/framework"
    google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
    tf "github.com/tensorflow/tensorflow/tensorflow/go"
    "google.golang.org/grpc"
)

type InceptionClient struct {
    ServingAddress string
}
// Predictor
func (i InceptionClient) Predict(imageBytes []byte) (*pb.PredictResponse, error) {
    tensor, err := tf.NewTensor(string(imageBytes))
    if err != nil {
        log.Fatalln("Cannot read image file")
        return nil, err
    }
    tensorString, ok := tensor.Value().(string)
    if !ok {
        log.Fatalln("Cannot type assert tensor value to string")
        return nil, err
    }
    request := &pb.PredictRequest{
        ModelSpec: &pb.ModelSpec{
            Name:          "inception",
            SignatureName: "predict_images",
            Version: &google_protobuf.Int64Value{
                Value: int64(1),
            },
        },
        Inputs: map[string]*tf_core_framework.TensorProto{
            "images": &tf_core_framework.TensorProto{
                Dtype: tf_core_framework.DataType_DT_STRING,
                TensorShape: &tf_core_framework.TensorShapeProto{
                    Dim: []*tf_core_framework.TensorShapeProto_Dim{
                        &tf_core_framework.TensorShapeProto_Dim{
                            Size: int64(1),
                        },
                    },
                },
                StringVal: [][]byte{[]byte(tensorString)},
            },
        },
    }
    log.Printf("serving address: %s", i.ServingAddress)
    conn, err := grpc.Dial(i.ServingAddress, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Cannot connect to the grpc server: %v\n", err)
        return nil, err
    }
    defer conn.Close()

    client := pb.NewPredictionServiceClient(conn)

    resp, err := client.Predict(context.Background(), request)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    return resp, nil
}
