<template>
  <div>
    <h1>Text to Speech</h1>
    <p>Enter text to convert to speech</p>
    <textarea></textarea>
    <button>Convert</button>
  </div>
  <ul v-for="voice in voliceList">
    <li>{{ voice.shortName }}</li>
  </ul>
</template>

<script>
  import { defineComponent } from 'vue'
  import axios from 'axios'

  export default defineComponent({
    data() {
      return {
        voliceList: []
      }
    },
    mounted() {
      axios.get('http://localhost:8888/list-voices/edge-tts', {
        withCredentials: true,  // 后端设置 Access-Control-Allow-Origin 为 * 时，前端不能设置 withCredentials
        timeout: 2000,
      })
        .then(response => {
          this.voliceList = response.data["voices"]
          console.log(response.data)
        })
        .catch(error => {
          console.log(error)
        })
    }
  })
</script>
