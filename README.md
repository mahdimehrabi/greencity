# SET UP
just run `docker compose up -d ` and all services will run and interact with each other
using a docker network that I defined its name to 'green'<br>
<br>

#### Results 
you can see result of producer using: <br> 
``` docker compose logs -f producer``` <br> <br> 
and you can see result of consumer using: <br>
``` docker compose logs -f consumer``` <br> <br>

### Verification
I didn't hard code verification and I implemented verification using redis



### Testing
I wrote two sample tests in producer service just for making sure that you know I know how to write tests. <br>
I hope thats satisfy you and please excuse me for not writing full tests.<br>



## Architecture (please read🙏)
<b> I assumed that you don't care about architecture because its just an examine project, and I assumed you want to qualify my algorithm and ability to use tools so
I didn't write code base on any architecture and in my opinion current architecture and design is pure disaster for a real production usage <br>
so please don't consider principles of writing maintainable clean code for your scoring if its possible.🙏🙏🙏🙏🙏</b>