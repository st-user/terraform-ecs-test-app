# terraform-ecs-test-app
App for testing codebuild with terraform

## Commands to build an image manually

```bash
ACCOUNT_ID=....
REGION=ap-northeast-1
ECR_IMAGE_TAG=${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com/fluentbit-dev-app:latest

aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com

docker build --platform linux/amd64 -t fluentbit-dev-app .

docker tag fluentbit-dev-app:latest ${ECR_IMAGE_TAG}

docker push ${ECR_IMAGE_TAG}

```
