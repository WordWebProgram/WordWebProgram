<template>
  <div class="home">
    <background />
    <Login v-if="showLogin" @close="showLogin=false" />
    <template v-if="user._id">
      <div class="main">
        <div class="revise">
          <div class="revise-left">
            <div class="revise-num">
              {{reviseNum}}
              <span>待复习</span>
            </div>
          </div>
          <div class="revise-right">
            <div class="revise-title">成为单词高手！</div>
            <div v-if="reviseNum" class="revise-button" @click="goRevise">开始复习</div>
          </div>
        </div>
        <div class="hint">新学单词：</div>
        <div class="learn">
          <div class="learn-list" @click="goLearn('list1')">
            <div class="list-title">Part 1</div>
            <div class="list-more">{{ getListHint('list1') }}</div>
            <div class="list-hover"></div>
          </div>
          <div class="learn-list" @click="goLearn('list2')">
            <div class="list-title">Part 2</div>
            <div class="list-more">{{ getListHint('list2') }}</div>
            <div class="list-hover"></div>
          </div>
          <div class="learn-list" @click="goLearn('list3')">
            <div class="list-title">Part 3</div>
            <div class="list-more">{{ getListHint('list3') }}</div>
            <div class="list-hover"></div>
          </div>
          <div class="learn-list" @click="goLearn('list4')">
            <div class="list-title">Part 4</div>
            <div class="list-more">{{ getListHint('list4') }}</div>
            <div class="list-hover"></div>
          </div>
          <div class="learn-list" @click="goLearn('list5')">
            <div class="list-title">Part 5</div>
            <div class="list-more">{{ getListHint('list5') }}</div>
            <div class="list-hover"></div>
          </div>
          <div class="learn-list" @click="goLearn('list6')">
            <div class="list-title">Part 6</div>
            <div class="list-more">{{ getListHint('list6') }}</div>
            <div class="list-hover"></div>
          </div>
        </div>
      </div>
    </template>
    <div v-else class="main">
      <div class="header">一起成为单词高手！</div>
      <img src="../static/icons/logo.svg" />
      <div class="start" @click="showLogin=true">现在开始</div>
    </div>

  </div>
</template>

<script>
import Login from '@/components/Login'
import word from '@/api/word'
import Background from '@/components/background.vue'

export default {
  name: 'home',
  components: {
    Background,
    Login
  },
  mounted () {
    this.$event.on('login', this, () => { this.showLogin = true })
    setTimeout(() => {
      this.getListProgress()
      this.getReviseNum()
    }, 200)
  },
  data () {
    return {
      showLogin: false,
      reviseNum: 0,
      listProgress: {}
    }
  },
  computed: {
    user () {
      return this.$store.state.user
    }
  },
  methods: {
    getListProgress () {
      word.getUserProgress().then((progress) => {
        this.listProgress = progress || {}
        console.log(progress)
      }).catch(err => console.log(err))
    },
    getReviseNum () {
      word.getReviseWordNum().then((num) => {
        this.reviseNum = num
      }).catch(err => console.log(err))
    },
    getListHint (listName) {
      const progress = this.listProgress[listName]
      const wordNum = word.getListWordNum(listName)
      if (progress) {
        const learnedNum = progress.location
        if (learnedNum >= wordNum) return '已完成学习'
        return `已学习${learnedNum}词`
      } else {
        return `尚未开始`
      }
    },
    goRevise () {
      this.$router.push('/revise')
    },
    goLearn (listName) {
      this.$router.push(`/learn?list=${listName}`)
    }
  }
}
</script>

<style scoped>
.home {
  width: 100vw;
  height: 100vh;
  background: #F0F0F0;
  overflow-x: hidden;
  overflow-y: auto;
  -webkit-overflow-scrolling: touch;
}

.main {
  height: fit-content;
  padding: 110px 30px;
  box-sizing: border-box;
  max-width: 1000px;
  margin: 0 auto;
}

.revise {
  width: 100%;
  background: #FFFFFF;
  box-shadow: 0 2px 7px 0 rgba(0, 0, 0, .24);
  border-radius: 4px;
  padding: 35px;
  box-sizing: border-box;
  display: flex;
  z-index: 1000;
  position:relative;
}

.revise-left {
  width: 140px;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.revise-num {
  width: 115px;
  height: 115px;
  border-radius: 50%;
  border: 9px solid #C1CACE;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  font-size: 25px;
  color: #40BC96;
  font-weight: 600;
  border: 8px solid #C1CACE;
}

.revise-num span {
  margin-top: 6px;
  font-size: 15px;
  color: #455358;
  font-weight: 700;
}

.revise-right {
  width: calc(100% - 140px);
  height: 100%;
  padding-left: 35px;
  box-sizing: border-box;
  z-index: 1000;
  background: #FFFFFF;
}

.revise-title {
  width: 100%;
  font-size: 36px;
  font-weight: 700;
  color: #455358;
  text-align: left;
}

.revise-hint {
  width: 100%;
  color: #455358;
  font-size: 17px;
  margin-top: 15px;
  text-align: left;
}

.revise-button {
  width: 220px;
  height: 64px;
  margin-top: 30px;
  background: #BD3124;
  color: #FFFFFF;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 17px;
  font-weight: 600;
  letter-spacing: 2px;
  cursor: pointer;
  transition: all 0.5s ease;
}

.revise-button:hover {
  background: #F9CC28;
  color: #455358;
}

.hint {
  width: 100%;
  font-weight: 600;
  font-size: 18px;
  color: #455358;
  margin-top: 30px;
  text-align: left;
}

.learn {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  margin-top: 15px;
}

.learn-list {
  width: 32%;
  border-radius: 5px;
  box-shadow: 0 0 1.5px 0 rgba(0,0,0,.24);
  background: #FFFFFF;
  padding: 20px 0 20px 20px;
  box-sizing: border-box;
  margin-bottom: 20px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.list-hover {
  width: 100%;
  height: 5px;
  background: #F9CC28;
  opacity: 0;
  transition: all 0.4s;
  position: absolute;
  bottom: 0;
  left: 0;
}

.learn-list:hover .list-hover {
  opacity: 1;
}

.list-title {
  width: 100%;
  font-size: 21px;
  font-weight: 700;
  color: #455358;
  text-align: left;
}

.list-hint {
  width: 100%;
  color: #99A5AA;
  font-size: 14px;
  margin-top: 8px;
  text-align: left;
  font-weight: 700;
}

.list-more {
  width: 100%;
  color: #455358;
  font-size: 13px;
  margin-top: 35px;
  text-align: left;
  font-weight: 600;
}

.header {
  width: 100%;
  color: #40BC96;
  text-align: center;
  font-size: 42px;
  letter-spacing: 3px;
  font-weight: bold;
  margin-top: 40px;
  margin-bottom: 80px;
}

.start {
  width: 240px;
  height: 64px;
  margin: 100px auto 0 auto;
  background: #3CCFCF;
  color: #FFFFFF;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 22px;
  font-weight: 600;
  letter-spacing: 2px;
  text-decoration: none;
  transition: all 0.5s ease;
  cursor: pointer;
}

.start:hover {
  background: #F9CC28;
  color: #455358;
}

.info {
  position: absolute;
  bottom: 0;
  height: 80px;
  width: 100%;
  background: #FFFFFF;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: 15px;
  font-weight: 500;
  color: #999999;
  line-height: 30px;
  cursor: pointer;
}

.info span {
  font-size: 14px;
}
</style>
