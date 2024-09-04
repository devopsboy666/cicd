# สร้าง App ด้วยภาษา Golang โดยจะมีทั้งหมดสาม Path คือ
1. / เป็น path แรกสุด
2. /p1
3. /status ที่เพื่อใช้ในการทำ Readiness Probes และ Liveness Probes

# สร้างของขึ้นเพื่อทำขั้นตอน CICD ดังนี้ 
1. สร้าง Dockerfile
2. สร้าง YAML File โดยจัดเก็บของทั้งหมดไว้ที่ Directory k8s - (Deployment, Service, Ingress)
3. สร้าง Jenkins File สำหรับ ทำ CICD

# เตรียม Cluster K8s ให้พร้อมเพื่อใช้งาน ดังนี้
1. สร้าง namespace ใหม่ขึ้นมาสำหรับ Deploy App
```
kubectl create ns app
```

2. สร้าง service account jenkins
```
kubectl create sa jenkins-robot
```

3. สร้าง RoleBinding ให้ Service Account Jenkins มีสิทธ์ในการสร้างแก้ไข้และลบ ที่ namesapce app
```
kubectl -n app create rolebinding jenkins-robot-binding --clusterrole=cluster-admin --serviceaccount=app:jenkins-robot
```

4. สร้าง Token แบบถาวรให้กลับ Jenkins
```
vi jenkins-robot.yaml
```
```
apiVersion: v1
kind: Secret
metadata:
  name: jenkins-robot
  annotations:
    kubernetes.io/service-account.name: jenkins-robot
```
```
kubectl create -f jenkins-robot.yaml
```
5. ดึง Token และ CA มาเก็บไว้เพื่อเตรียมทำไฟล์ kubeconfig เฉพาะของ Service Account Jenkins
```
# ดึง token
kubectl get secret jenkins-robot -o jsonpath='{.data.token}' -n app | base64 -d
# ดึง CA
kubectl get secret jenkins-robot -o jsonpath='{.data.ca\.crt}' -n app
```
6. นำค่า Token กับ CA ที่ได้มาทำการสร้างไฟล์ kubeconfig 
```
vi config
```
แก้ไข 3 จุดคือ certificate-authority-data, server, token หลังแก้เสร็จนำไฟล์ไปเตรียมไว้ ที่ Jenkins Server เราสามารถเข้าถึงได้เพื่อทำการ Import
```
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: <CA>
    server: https://<api-kubernetes-server>:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    namespace: app
    user: jenkins-robot
  name: jenkins-robot
current-context: jenkins-robot
kind: Config
preferences: {}
users:
- name: jenkins-robot
  user:
    token: <TOKEN>
```
7. สร้าง image pull secret สำหรับ nexus
```
kubectl -n app create secret docker-registry nexus-pull-secret --docker-server=<nexus_server> --docker-username=<username> --docker-password=<password>
```

# เพิ่ม Credentials nexus และ kubeconfig 
1. ลง Plugins เพิ่มเติมดังนี้
```
Dashboard -> Manage Jenkins -> Plugins -> Avaliable Plugins -> Search "Pulgin Docker pipeline"
Dashboard -> Manage Jenkins -> Plugins -> Avaliable Plugins -> Search "Pulgin kubernetes CLI"
Dashboard -> Manage Jenkins -> Plugins -> Avaliable Plugins -> Search "Pulgin Stage View"
```
2. เพิ่ม credentials nexus สำหรับ push และ pull image
```
Dashboard -> Manage Jenkins -> Credentials -> System -> Global credentials (unrestricted) -> Add Credentails -> กรอกข้อมูลลงไป โดยเลือก Kind = Username with password (ในตัวอย่างตั้งชื่อว่า nexus-credentials)
```
3. เพิ่ม credentials ของ kubeconfig
```
Dashboard -> Manage Jenkins -> Credentials -> System -> Global credentials (unrestricted) -> Add Credentails -> กรอกข้อมูลลงไป โดยเลือก Kind = Secret file (ในตัวอย่างตั้งชื่อว่า context)
```
4. สร้าง pipeline ขึ้นมา
```
Dashboard -> New Item -> เลือก Pipeline -> Pipeline Definition เลือกเป็น "Pipeline script from SCM"
```
