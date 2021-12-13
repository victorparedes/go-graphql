# Go + GraphQL como servidor
Ejemplo practico de una implementacion de GraphQL con Go.


Desrrollado con Go v1.17.5

```console
git clone https://github.com/victorparedes/nodejs-graphql.git
cd nodejs-graphql
npm install
node index.js 
```
Luego accedemos a http://localhost:5000/graphql
<br/><br/>

## El problema con las RESTAPIs y modelos de datos estaticos.

Imaginemos que somos una Red Social (como Facebook) y necesitamos poder mostrar el perfil de un usuario en un telefono o en una pagina web. En la pagina web mostrariamos todo los datos del usuario ya que tenemos lugar en pantalla... Peeeeero, si estamos en un telefono no tenemos mucho espacio disponible asi que solo mostramos los datos minimos. Si usaramos una Rest Api tradicional nos encontrariamos con los siguientes problemas:

**- Solicitariamos informacion de mas:** Si usamos el mismo endpoint para ambos casos, en el telefono traeremos informacion que no usaremos (ya que no vamos a poder mostrar todo en pantalla), consumiremos datos y bateria posicionando nuesta APP en los primeros puestos de consumo de estos preciados recursos (con lo cual, si no somos facebook, nos arriesgamos a que nos desinstalen).

**- Deberiamos tener dos endpoint para cada dispositivo:** Si hacemos dos endpoints (uno para web y otro para telefono) deberemos mantener ambos funcionando. Tendremos problemas si deseamos expandirnos a otras plataformas (tablets o smartTVs). Incluso, si reutilizamos el codigo fuente para ambos endpoint, estariamos obligando a nuestro servidor a calcular toda la informacion de manera innecesaria.

<br/>
Ohh.. y ahora... ¿quien podra ayudarnos?

<br/>

**CUIDADO!!! GrapqhQL no es mejor que REST API ni REST API es mejor que GraphQL. Ambas son herramientas muy poderosas que sirven para resolver un problema especifico en cada caso (GraphQL es un martillo, Rest Api es un destornillador).**

# Resolviendo problemas con consultas en tiempo de ejecucion (query API)

Para resolver el problema presentado Facebook creo GraphQL. Esta herramienta nos permite ejecutar *querys* a una API como si lo hicieramos contra una base de datos relacional (SQL)

GraphQL nos permite exponer nuestra API de tal forma que podamos consumirla con una query (si... parecido a SQL) en tiempo de ejecucion. Si quisiera podria pedir uno o todos los campos del perfil de un usuario (o la combinacion de ellos) a un solo endpoint. Podria agregar propiedades a los usuarios sin enviar esa informacion nueva a procesos que no la necesitan.

```
// Ejemplo de una query que me traeria informacion para un telefono
// trae unos pocos campos y solo las imagenes para mobile.
{
    UserProfile(email: 'example@gmail.com') {
        name
        lastName
        mobileProfileImage
        haveNotification
        lastPublications {
            mobileImage
            title
        }
    }
}
```

```
// Ejemplo de una query que traeria informacion para la web
// trae todos los campos y solo las imagenes para WEB
{
    UserProfile(email: 'example@gmail.com') {
        name
        lastName
        profileDescription
        webProfileImage
        friends {
            webImage
            name
            lastName
            profileUrl
        }
        lastPublications {
            webImage
            title
        }
        photos {
            webImage
            title
        }
        notifications{
            title
            notificationType
        }
    }
}
```

# Principales ventajas

- Es muy facil de usar y aprender. Cuando se logra entender el concepto y como funciona lleva muy poco tiempo escribir un servidor productivo. Implementado correctamente se logra dividir los componentes de manera tan atomica que permite realizar TDD con casos muy simples.

- Soluciona el overfitting y el underfitting que es el problema de traer demasiados datos o de traer pocos (lo que necesitaria una nueva llamada a otro servicio).

- Optimiza el uso de la red y los datos ya que permite ejecutar varias querys en el mismo request. Es decir que podriamos pedir datos de un usuario, stock actual de un producto y los precios segun las diferentes listas con una sola consulta al servidor.

- Puede obtener informacion de mas de un proveedor de datos. Es decir que podemos obtener datos desde cualquier servicio, base de datos o repositorio. Por ejemplo, si una query pide datos usuarios, podemos obtenerlos desde SQL, luego consultar al Active Directory para conocer los permisos y finalmente buscar la imagen del perfil en una [CDN](https://es.wikipedia.org/wiki/Red_de_distribuci%C3%B3n_de_contenidos) y devolverlo a quien lo pidio.

- El servidor solo procesara la informacion que se le pide. Si por algun motivo (siguiendo con el ejemplo anterior del usuario) no necesito conocer los permisos que el usuario tiene la consulta al Active Directory no se realizara. Lo mismo sucederia con la CDN si no pido la imagen de perfil.

# Componentes basicos

Un servidor GraphQL esta compuesto por, al menos, cuatro componentes principales y necesarios.

## Servidor
GrapqhQL por si solo no puede funcionar como un servicio web, requiere de algun componente que le se soporte HTTP (como express, restify y HTTP nativo). En nuestro caso usaremos [Express](https://www.npmjs.com/package/express) junto con el modulo [Express-GraphQL](https://www.npmjs.com/package/express-graphql).

## Schema
Un *schema*, en palabras simples, es la lista de *querys* que nuestro servidor GraphQL va a soportar. Por ejemplo, en este codigo, tenemos cuatro *querys*: product, products, description y modality.

Un schema esta compuesto de tres elementos:

**- Type:** Un tipo de dato a devolver (o una lista de ellos)

**- Arguments:** Parametros que se especifican en la *query* y que se utilizan para filtrar.

**- Resolvers:** La implementacion del codigo necesario para devolver el type que se le asigno.

## Types
Los *types* son los datos que devolvera una *query*. Estos datos estan compuestos de propiedades fuertemente tipadas y perfectamente definidas.
Cada propiedad tiene un *resolver* (responsable de devolver el dato) que puede estar implicito por defecto o customizado por el desarrollador segun se necesite.

## Resolvers
Los *resolvers* son basicamente el codigo necesario para devolver el valor que se asignara a una propiedad. Un resolver puede ser implicito, es decir, que se devolvera el dato que tenga el mismo nombre que la propiedad o, puede ser personalizado para tomar un ID de los datos padres y buscar los valores a devolver desde otro origen de datos.

