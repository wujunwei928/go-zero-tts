import { get, postJson } from "@/utils/request";

export const getVoiceList = (platform) => get(`/list-voices/${platform}`)

export const textToSpeech = (platform, voice, text, volume, pitch, rate) => postJson("/text-to-speech", { platform, voice, text, volume, pitch, rate})