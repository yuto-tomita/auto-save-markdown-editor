import { defineStore } from 'pinia';
import { ApiState } from './type';
import axios from 'axios';

export const apiStore = defineStore({
	id: 'markdownApi',
	state: (): ApiState => ({ DraftList: {} }),
	getters: {
    returnCalendarList (state): ApiState["DraftList"]  {
      return state.DraftList;
    }
  },
  actions: {
    async getDraftList (): Promise<void> {
      try {
				console.log('come');
        const res = await axios.get(import.meta.env.VITE_API_URI);
				console.log(res);

				this.DraftList = res;
      } catch (e) {
        console.log(e);
      }
    }
  }
});