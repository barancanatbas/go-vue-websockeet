<template>
  <div class="container">
    <div class="chat__body" id="chat__body">
        <div v-for="item in messages" v-bind:key="item.id">

          <div v-if="item.sender == 1"> <p class="chat__mymessage__paragraph">{{item.greeting}}</p> </div>

          <div v-else>
            <div>
              <div class="chat__yourmessage__p">
                <p class="chat__yourmessage__paragraph">
                  {{item.greeting}}
                </p>
              </div>
            </div>
          </div>

        </div>
      </div>
      <input type="text" class="msg" v-model="message">
      <input type="submit" class="send" value="Gonder" v-on:click="sendMessage">
      
  </div>
</template>

<script>

export default {
  name: 'App',
  data(){
    return {
      message:"",
      socket:null,
      messages: []
    }
  },
  mounted(){
    this.socket = new WebSocket("ws://localhost:9100/socket")
    this.socket.onmessage = (msg) =>{
      console.log(JSON.parse(msg.data))
      this.messages.push(JSON.parse(msg.data))
    }
  },
  methods:{
    sendMessage(){
      let msg = {
        "greeting":this.message
      }
      this.socket.send(JSON.stringify(msg))

      this.messages.push({greeting:this.message,sender:1})
      this.message = ""
    }
  }
}
</script>

<style>
.container{
  width: 500px;
  height: auto;
  margin: 200px auto;
}
.send{
  padding: 12px 20px;
  border: none;
  margin-left: 20px;
}
.msg{
  width: 300px;
  padding: 10px 20px;
  border: none;
  box-shadow: 1px 0px 62px -2px rgba(0,0,0,0.10);
-webkit-box-shadow: 1px 0px 62px -2px rgba(0,0,0,0.10);
-moz-box-shadow: 1px 0px 62px -2px rgba(0,0,0,0.10);
}
 .chat__body {
   width: 500px;
  padding: 2rem;
  overflow: scroll;
  scroll-behavior: smooth;
}
.chat__body::-webkit-scrollbar {
  display: none;
}
.chat__mymessage {
  display: flex;
  justify-content: right;
  align-items: flex-end;
  margin: 0;
  min-height: 40px;
  line-break: anywhere;
}
.chat__mymessage__paragraph {
  margin: 0.4rem 0 0 1rem;
  border-radius: 20px 20px 0px 20px;
  max-width: 220px;
  background-color: #bbc4ef;
  color: #ffffff;
  padding: 0.8rem;
  font-size: 14px;
}
.chat__first {
  margin-top: 2rem;
}
.chat__yourmessage {
  display: flex;
}

.chat__yourmessage__user {
  font-size: 20px;
  font-weight: normal;
  color: #292929;
  margin-top: 0;
  margin-block-end: 0rem;
}
.chat__yourmessage__p {
  display: flex;
  align-items: flex-end;
  line-break: anywhere;
}
.chat__yourmessage__paragraph {
  margin: 0.4rem 1rem 0 0;
  border-radius: 0px 20px 20px 20px;
  background-color: #f3f3f3;
  max-width: 220px;
  color: #414141;
  padding: 0.8rem;
  font-size: 14px;
}
.chat__yourmessage__time {
  margin: 0;
  font-size: 12px;
  color: #9c9c9c;
}
</style>