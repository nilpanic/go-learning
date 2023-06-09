<script setup>
import { NButton, NH2, NInput, NSpace, useMessage } from "naive-ui";
import { ref } from "vue";
import http from "@/lib/http";

const message = useMessage()

const qrUrl = ref('')
const qrImg = ref('')
const showQr = ref(false)
const loading = ref(false)

function createQrCode () {
  loading.value = true
  if (qrUrl.value === '') {
    message.error('请输入url地址')
    loading.value = false
    return
  }

  http.post('http://localhost:3000/v1/qrcode', { url: qrUrl.value }).then(async res => {
    qrImg.value = res.data.data.url;
    showQr.value = true
  }).catch(err => {
    err = err ? err.message : '网络错误'
    message.error(err)
  }).finally(() => {
    loading.value = false
  })
}

</script>

<template>
  <div class="qrcode-box">
    <n-h2>二维码</n-h2>
    <n-space>
      <n-input style="width: 300px" v-model:value="qrUrl" type="text" placeholder="请输入url地址"/>
      <n-button @click="createQrCode" :loading="loading" type="primary" style="width: 90px">创建</n-button>
      <Transition name="fade">
        <div class="qrcode-rel" v-if="showQr">
          <img :src="qrImg" alt=""/>
          <n-button style="width: 200px" type="primary" strong secondary>立即下载</n-button>
        </div>
      </Transition>
    </n-space>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
.qrcode-box {
  padding: 20px 0;
  width: 800px;
  margin: 0 auto;
}

.qrcode-rel {
  display: inline-block;
  width: 200px;
}

.qrcode-rel img {
  width: 200px;
  height: 200px;
  border: 1px solid #ddd;
  border-radius: 3px;
  display: inline-block;
}

.n-h2 {
  margin-top: 50px;
}
</style>
