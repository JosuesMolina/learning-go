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
