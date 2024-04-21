<template>
  <v-container fluid>
    <v-autocomplete
      v-model="selectVoice"
      :items="voliceList"
      item-title="shortName"
      item-value="shortName"
      label="选择声音"
      v-com
      return-object
    ></v-autocomplete>
  </v-container>
  <v-container fluid>
    <v-card style="margin: auto">
      <v-card-text>
        <v-slider
          v-model="volume"
          :min="0.5"
          :max="2"
          :step="0.1"
          class="ma-4"
          label="音高"
          hide-details
        >
          <template v-slot:append>
            <v-text-field
              v-model="volume"
              density="compact"
              style="width: 80px"
              type="number"
              variant="outlined"
              readonly
              hide-details
            ></v-text-field>
          </template>
        </v-slider>

        <v-slider
          v-model="pitch"
          :min="0.5"
          :max="2"
          :step="0.1"
          class="ma-4"
          label="语速"
          hide-details
        >
          <template v-slot:append>
            <v-text-field
              v-model="pitch"
              density="compact"
              style="width: 80px"
              type="number"
              variant="outlined"
              readonly
              hide-details
            ></v-text-field>
          </template>
        </v-slider>

        <v-slider
          v-model="rate"
          :min="0.5"
          :max="2"
          :step="0.1"
          class="ma-4"
          label="音量"
          hide-details
        >
          <template v-slot:append>
            <v-text-field
              v-model="rate"
              density="compact"
              style="width: 80px"
              type="number"
              variant="outlined"
              readonly
              hide-details
            ></v-text-field>
          </template>
        </v-slider>
      </v-card-text>
    </v-card>
  </v-container>
  <v-container fluid>
    <v-textarea
      label="输入文本"
      v-model="inputText"
      variant="filled"
      auto-grow
    ></v-textarea>
  </v-container>

  <v-container fluid>
    <v-btn
      @click="textToSpeech"
      rounded="xl"
      size="x-large"
      block
    >生成声音</v-btn>
  </v-container>

  <v-container :style="{'display': mp3ContainerDisplay}">
    <h1>播放 文本转语音 MP3 文件内容</h1>
    <audio controls id="audioPlayer" :src="mp3Url">
          Your browser does not support the audio element.
    </audio>
  </v-container>
</template>

<script>
  import { defineComponent } from 'vue'
  import axios from 'axios'

  export default defineComponent({
    data() {
      return {
        pitch: 1,
        rate: 1,
        volume: 1,
        inputText: "白日依山尽,黄河入海流.欲穷千里目,更上一层楼.", //白日依山尽,黄河入海流.欲穷千里目,更上一层楼.
        selectVoice: {},
        voliceList: [],
        mp3ContainerDisplay: "none",
        mp3Url: ""
      }
    },

    methods: {
      textToSpeech() {
        this.mp3ContainerDisplay = "none";
        this.mp3Url = "";

        const selectVoice = this.selectVoice;
        const inputText = this.inputText;
        console.log(selectVoice)
        console.log(inputText)
        const postData = {
          platform: "edge-tts",
          voice: this.selectVoice.shortName,
          text: this.inputText,
        }

        axios.post(
          'http://localhost:8888/text-to-speech',
          postData,
          {
            headers: {
              'Content-Type': 'application/json'
            },
            timeout: 10000,
          }
        )
        .then(response => {
          // this.voliceList = response.data["voices"]
          // this.selectVoice = this.voliceList[0]
          
          // 将 base64 编码的音频数据解码成二进制数据
          let binaryAudioData = atob(response.data.audio);

          // 将二进制数据转换成 Uint8Array
          let byteArray = new Uint8Array(binaryAudioData.length);
          for (let i = 0; i < binaryAudioData.length; i++) {
              byteArray[i] = binaryAudioData.charCodeAt(i);
          }

          // 创建 Blob 对象
          let blob = new Blob([byteArray], { type: 'audio/mpeg' });

          // 创建 URL 对象
          let audioURL = URL.createObjectURL(blob);

          // 将音频文件 URL 赋给 audio 元素的 src 属性
          this.mp3Url = audioURL;
          this.mp3ContainerDisplay = "block";
        })
        .catch(error => {
          console.log(error)
        })
      },
    },

    mounted() {
      axios.get('http://localhost:8888/list-voices/edge-tts', {
        withCredentials: true,  // 后端设置 Access-Control-Allow-Origin 为 * 时，前端不能设置 withCredentials
        timeout: 2000,
      })
        .then(response => {
          this.voliceList = response.data["voices"]
          this.selectVoice = this.voliceList[0]
          console.log(response.data)
        })
        .catch(error => {
          console.log(error)
        })
    }
  })
</script>
