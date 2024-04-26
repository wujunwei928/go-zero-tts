<template>
  <v-container fluid>
    <v-autocomplete
      v-model="selectLocale"
      :items="localeList"
      item-title="displayName"
      item-value="shortName"
      label="选择地区"
      @update:model-value="changeCountry($event)"
      v-com
      return-object
    ></v-autocomplete>
    <v-autocomplete
      v-model="selectVoice"
      :items="countryVoiceList"
      item-title="displayName"
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
          :min="0.1"
          :max="3"
          :step="0.1"
          class="ma-4"
          label="音量"
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
          :min="0.1"
          :max="3"
          :step="0.1"
          class="ma-4"
          label="音高"
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
          :min="0.1"
          :max="3"
          :step="0.1"
          class="ma-4"
          label="语速"
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
      @click="textToSpeechApi"
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

  <!-- 错误信息 -->
  <v-alert
   type="error"
   v-show="submitError"
  >
    {{ submitErrorMessage }}
  </v-alert>
</template>

<script setup>
  import { ref, computed, onMounted } from 'vue'
  import axios from 'axios'
  import { useLocalStorage } from '@vueuse/core';
  import { getVoiceList, textToSpeech } from '@/api/tts';
  import { cnLocalLanguage, cnLocalVoice } from '@/utils/voiceConfig';

  let pitch = ref(1)
  let rate = ref(1)
  let volume = ref(1)
  let inputText = ref("白日依山尽,黄河入海流.欲穷千里目,更上一层楼.")
  let mp3ContainerDisplay = ref("none")
  let mp3Url = ref("")
  let submitError = ref(false)
  let submitErrorMessage = ref("")
  // 使用useLocalStorage获取或初始化声音列表
  let voiceList = useLocalStorage('voiceList', [])
  let countryVoicesMap = ref({})
  // 选择地区和声音
  let selectLocale = ref(undefined)
  let selectVoice = ref(undefined)
  let localeList = ref([])
  let countryVoiceList = ref([])
  

  const textToSpeechApi = async () => {
    try {
      // 错误信息置空
      submitError.value = false
      submitErrorMessage.value = ""

      // 清空音频文件
      mp3Url.value = "";
      mp3ContainerDisplay.value = "none";

      if (selectVoice.value === undefined) {
        submitError.value = true
        submitErrorMessage.value = "请选择声音"
        return
      }
      
      const platform = "edge-tts"
      const voice = selectVoice.value.shortName
      const text = inputText.value
      const volumeValue = parseFloat(volume.value)
      const pitchValue = parseFloat(pitch.value)
      const rateValue = parseFloat(rate.value)

      const response = await textToSpeech(platform, voice, text, volumeValue, pitchValue, rateValue)
      
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
      mp3Url.value = audioURL;
      mp3ContainerDisplay.value = "block";
    } catch (error) {
      submitError.value = true
      submitErrorMessage.value = error.response.data
    }
  }

  function changeCountry(event) {
    doChangeCountry(event.shortName)
  }

  // 切换地区时触发
  const doChangeCountry = (country) => {
    // 默认选择 shortName 为 zh-CN 的地区
    selectLocale.value = {
      displayName: cnLocalLanguage[country],
      shortName: country
    }

    // 获取地区对应的声音列表
    countryVoiceList.value = countryVoicesMap.value[selectLocale.value.shortName].map((item) => {
      return {
        displayName: cnLocalVoice[item.shortName],
        shortName: item.shortName
      }
    })

    // 默认选择第一个声音
    selectVoice.value = countryVoiceList.value[0]
  }

  const getVoiceListApi = async () => {
    try {
      const response = await getVoiceList("edge-tts")
      voiceList.value = response.data["voices"]
    } catch (error) {
      console.log(error)
    }
  }
  onMounted(() => {
    // 当组件挂载时检查localStorage，如果没有声音列表则发起请求
    if (voiceList.value.length === 0) {
      getVoiceListApi()
    }

    // 将声音列表按地区分类
    countryVoicesMap.value = voiceList.value.reduce((acc, cur) => {
      if (cur.locale in acc) {
        acc[cur.locale].push(cur)
      } else {
        acc[cur.locale] = [cur]
      }
      return acc
    }, {})

    // 获取声音地区列表, 获取 countryVoicesMap 的key
    localeList.value = Object.keys(countryVoicesMap.value).map((item) => {
      return {
        displayName: cnLocalLanguage[item],
        shortName: item
      }
    })

    // 默认选择 中国 zh-CN
    doChangeCountry("zh-CN")
  })

</script>
