import { apiCall } from './framework.js';
import { Info } from '../ts-gen/info.js'

async function getSeekrInfo(): Promise<Info> {
  const response = await fetch(apiCall('/info'));
  const data = await response.json();
  return data as Info;
}

export { getSeekrInfo };
