import { apiCall } from './framework.js';

interface SeekrInfo {
    download_url: string;
    is_latest:    boolean;
    latest:       string;
    version:      string;
}

async function getSeekrInfo(): Promise<SeekrInfo> {
  const response = await fetch(apiCall('/info'));
  const data = await response.json();
  return data as SeekrInfo;
}

export { SeekrInfo, getSeekrInfo };
