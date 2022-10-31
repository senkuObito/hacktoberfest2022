resource "aws_athena_data_catalog" "example" {
  name        = "athena-data-catalog"
  description = "Example Athena data catalog"
  type        = "LAMBDA"

  parameters = {
    "function" = "arn:aws:lambda:eu-central-1:123456789012:function:not-important-lambda-function"
  }

  tags = {
    Name = "example-athena-data-catalog"
  }
}
