import { defineStore } from 'pinia';
import { ApiState, DraftParameters } from './type';
import axios from 'axios';

export const apiStore = defineStore({
	id: 'markdownApi',
	state: (): ApiState => ({ DraftList: [] }),
	getters: {
    returnCalendarList (state): ApiState["DraftList"]  {
      return state.DraftList;
    }
  },
  actions: {
    async getDraftList (): Promise<void> {
      try {
        const res = await axios.get(import.meta.env.VITE_API_URI);

				this.DraftList = res.data.res;
      } catch (e) {
        console.log(e);
      }
    },
    async saveMarkdownDraft (params: DraftParameters): Promise<void> {
      try {
        await axios.post(`${import.meta.env.VITE_API_URI}/saveMarkDown`, params);
      } catch (e) {
        console.log(e);
      }
    },
    async deleteDraft (draftId: number): Promise<void> {
      try {
        await axios.delete(`${import.meta.env.VITE_API_URI}/draft/${draftId}`);
      } catch (e) {
        console.log(e);
      }
    }
  }
});