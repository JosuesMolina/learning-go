# El siguiente proyecto tiene como finalidad recopilar y exponer notas relevantes de mi aprendizaje en GO.

## Compilación

- ``go build``

Crea compilado con un nuevo nombre.  
- ``go build -o new-name main.go``

agregar la ruta del proyecto al PATH  
1. busca el archivo ``.profile ``
- ``nano .profile``  
2. agrega la ruta:  
- ``export PATH=$PATH:/home/user/go/src/proyect-name``
3. actualizar los cambios en el file .profile  
- ``source .profile``

1. compilar e instalar dentro de la carpeta bin  
- ``go install``

ejecutar app go and wait response  
- ``name-app.exe && read -p "press to continue" ``

## Iniciales    
``go env`` => ver todas las variables de GO

``go.mod`` => define y gestiona las dependencias
- example: ``go mod init name-proyect``
  
1. Agregar una dependencia:
- ``go mod edit -require=example.com/somepackage@v1.2.3``
﻿
2. Eliminar una dependencia:
- ``go mod edit -droprequire=example.com/somepackage``

3. Cambiar el nombre del módulo:
- ``go mod edit -module=example.com/newname``

4. Reemplazar una dependencia temporalmente:
- ``go mod edit -replace=example.com/somepackage=../path/to/local/package``

5. ``go mod tidy``
- Eliminar dependencias innecesarias que ya no se utilizan en el código.  
- Añadir dependencias faltantes que se están utilizando pero no están listadas.  
- Mantener el archivo go.mod y go.sum organizados y actualizados, lo que facilita la gestión de las dependencias en un proyecto de Go.
