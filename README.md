# Prueba Técnica para Ingeniero DevOps Junior

## Objetivo

Tu tarea es demostrar tus habilidades en DevOps creando una infraestructura en **AWS** utilizando **Terraform** y configurando un **pipeline CI/CD** con **GitHub Actions**. El objetivo es desplegar una instancia EC2, configurar una red y crear una infraestructura escalable que ejecute una aplicación proporcionada.

**Nota**: Para todas las tareas, utiliza únicamente la documentación oficial de **AWS**, **Terraform**, **Docker** y **GitHub Actions**.

---

## Resumen de Instrucciones

Deberás crear:

1. Una **instancia EC2** en AWS que ejecute una aplicación Dockerizada desde el repositorio de GitHub proporcionado.
2. Una configuración de red completa utilizando **Terraform**, que incluya:
    - **VPC** con subredes públicas y privadas.
    - Un **Auto Scaling Group** y un **Load Balancer**.
    - **Certificados SSL/TLS** utilizando AWS Certificate Manager.
3. Un **pipeline CI/CD** con **GitHub Actions** para automatizar el proceso de despliegue de la infraestructura.

---

## Requisitos Técnicos

### 1. Configuración de EC2 y Docker

- **Crear** una instancia EC2 en AWS con los permisos necesarios y ejecutar el repositorio https://github.com/cloud-craftman/tech-test.
- **Instalar Docker** en la instancia EC2 para ejecutar la aplicación. Asegúrate de seguir la guía oficial de instalación de Docker para la distribución de Linux elegida.
- **Recursos de Docker**:
    - Guía de instalación de Docker: [how-to-install-docker-on-amazon-linux-2](https://www.cyberciti.biz/faq/how-to-install-docker-on-amazon-linux-2/)

### 2. Infraestructura como Código con Terraform

- **Crear** los siguientes recursos en AWS usando Terraform:
    - **VPC**:
        - Diseñar una VPC con subredes públicas y privadas.
        - Configurar un Internet Gateway y un NAT Gateway según sea necesario.
    - **Auto Scaling Group (ASG)**:
        - Configurar un Auto Scaling Group para administrar las instancias EC2.
        - Configurar una plantilla de lanzamiento con grupos de seguridad adecuados.
    - **Load Balancer (ALB)**:
        - Configurar un Application Load Balancer (ALB) para distribuir el tráfico.
        - Asegurarse de que funcione con los certificados SSL/TLS de AWS Certificate Manager.
    - **Certificados SSL/TLS**:
        - Usar AWS Certificate Manager para aprovisionar certificados para conexiones seguras.
- **Recursos**:
    - AWS: [Documentación de AWS](https://docs.aws.amazon.com/)
    - Terraform: [Documentación de Terraform](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

### 3. Pipeline CI/CD con GitHub Actions

Crear un **workflow de GitHub Actions** para despliegue continuo. El pipeline debe incluir los siguientes pasos:

1. **Terraform Init**: Inicializar el directorio de trabajo de Terraform.
2. **Terraform Validate**: Verificar si el código de Terraform es sintácticamente válido.
3. **Terraform Plan**: Generar un plan de ejecución para previsualizar los cambios.
4. **Terraform Apply**: Aplicar los cambios para configurar la infraestructura.

- **Recursos**:
    - GitHub Actions: [Documentación de GitHub Actions](https://docs.github.com/en/actions)

---

## Entregables

1. **Código de Terraform**: Un conjunto completo de archivos `.tf` que configuren la infraestructura de AWS.
2. **Workflow de GitHub Actions**: Un archivo `.yml` que defina el pipeline CI/CD.
3. **Documentación**:
    - Pasos detallados que expliquen cómo ejecutar y verificar la configuración.
    - Cualquier desafío que enfrentaste y cómo lo resolviste.

---

## Criterios de Evaluación

- **Corrección**: ¿Funciona la configuración como se espera?
- **Calidad del Código**: ¿Está el código de Terraform bien estructurado y modular?
- **Documentación**: ¿Es fácil de entender y replicar la configuración?
- **Resolución de Problemas**: ¿Cómo resolviste los problemas encontrados durante la configuración?
- **Uso de Documentación Oficial**: Asegúrate de referenciar y aprender de las fuentes oficiales.

---

## Puntos Extra

- Uso de **módulos de Terraform** para organizar el código.
- Configuración de **monitorización con CloudWatch** para tu infraestructura.
- Creación de **grupos de seguridad adicionales** con reglas restrictivas para mayor seguridad.

---

## Entrega

1. Haz un fork del repositorio de GitHub proporcionado y sube tus cambios.
2. Envía un enlace a tu repositorio de GitHub y un breve resumen de tu enfoque.

¡Buena suerte y feliz codificación! Recuerda usar la documentación oficial para guiarte durante todo el proceso.
