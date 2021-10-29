<template>
  <div class="main-content">
    <div class="content-left">
      <div class="content-left-item">
        <h4>企業情報の設定</h4>
        <hr class="line" />
        <div class="infomations">
          <div class="infomation">
            <label>企業のロゴ</label>
            <img
              class="image-info"
              :src="require('../assets/' + logo)"
              alt=""
            />
            <button class="edit" @click="showImage = true">
              <img src="../assets/edit.png" alt="" /><label>編集</label>
            </button>
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showImage"
                @click="showImage = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-photo" v-if="showImage">
                <div class="title" @click="showImage = false">
                  <h3>企業ロゴを変更</h3>
                  <button><img src="../assets/close.png" alt="" /></button>
                </div>
                <div class="images">
                  <img :src="require('../assets/' + logo)" alt="" />
                  <div class="upload-image">
                    <img src="../assets/icon_upload.png" alt="" />
                    <p>画像のアップロード</p>
                  </div>
                </div>
                <button class="update">変更する</button>
              </div>
            </transition>
          </div>
          <div class="infomation">
            <label>企業名</label>
            <p>{{ company_name }}</p>
            <button class="edit" @click="showName = true">
              <img src="../assets/edit.png" alt="" /><label>編集</label>
            </button>
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showName"
                @click="showName = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-name" v-if="showName">
                <div class="title" @click="showName = false">
                  <h3>企業ロゴを変更</h3>
                  <button><img src="../assets/close.png" alt="" /></button>
                </div>
                <div class="input-name">
                  <label for="">企業名</label>
                  <input value="〇〇株式会社" type="text" />
                </div>
                <button class="update">変更する</button>
              </div>
            </transition>
          </div>
          <div class="infomation">
            <label>企業住所</label>
            <p>{{ address_company }}</p>
            <button class="edit">
              <img src="../assets/edit.png" alt="" /><label>編集</label>
            </button>
          </div>
          <div class="infomation">
            <label>電話番号</label>
            <p>{{ phone }}</p>
            <button class="edit">
              <img src="../assets/edit.png" alt="" /><label>編集</label>
            </button>
          </div>
          <div class="infomation">
            <label>代表者氏名</label>
            <p>{{ name }}</p>
            <button class="edit">
              <img src="../assets/edit.png" alt="" /><label>編集</label>
            </button>
          </div>
          <div class="infomation">
            <label>パートナー会社</label>
            <p>{{ co_partner }}</p>
          </div>
          <div class="infomation">
            <label>紹介営業担当者</label>
            <p>{{ seller_name }}</p>
          </div>
          <div class="infomation">
            <label>管理番号</label>
            <p>{{ check_number }}</p>
          </div>
          <div class="infomation">
            <label>オリジナル印影</label>
            <img
              class="image-info"
              :src="require('../assets/' + originer_imprint)"
              alt=""
            />
            <button class="edit">
              <img src="../assets/edit.png" alt="" /><label>編集</label>
            </button>
          </div>
        </div>
      </div>
      <div class="content-left-item">
        <h4>セキュリティ設定</h4>
        <hr class="line" />
        <div class="infomations">
          <div class="infomation">
            <label>
              <p>すべてのメンバーに</p>
              <p>二段階認証を要求する</p>
            </label>
            <p>オフ</p>
            <img
              class="toggle"
              @click="onClick1(security)"
              :src="
                security
                  ? require('../assets/toggle-on.png')
                  : require('../assets/toggle-off.png')
              "
              alt=""
            />
          </div>
        </div>
      </div>
      <div class="content-left-item">
        <h4>メール通知設定</h4>
        <hr class="line" />
        <div class="infomations">
          <div class="infomation">
            <label>メール通知を許可する</label>
            <p>オン</p>
            <img
              class="toggle"
              @click="onClick2(notification)"
              :src="
                notification
                  ? require('../assets/toggle-on.png')
                  : require('../assets/toggle-off.png')
              "
              alt=""
            />
          </div>
        </div>
      </div>
      <div class="add-content">
        <h4>アラート通知先の設定</h4>
        <button class="add-content-submit">
          <img src="../assets/icon_add.png" alt="" /><label
            >アラート通知先の設定をする</label
          >
        </button>
      </div>
      <hr class="line" />
      <footer class="footer-content">
        <p>アラート通知先のメールアドレスをデフォルト設定できます。</p>
        <p>また、他のメニューから契約書ごとに変更することもできます。</p>
      </footer>
    </div>
    <div class="content-right">
      <div class="content-right-item active"><a>企業情報の設定 </a></div>
      <div class="content-right-item"><a>セキュリティ設定 </a></div>
      <div class="content-right-item"><a>メール通知設定 </a></div>
      <div class="content-right-item"><a>アラート通知先の設定 </a></div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      showImage: false,
      showName: false,
      id: this.$route.params.id,
      accounts: {},
      logo: "", company_name: "", address_company: "", phone: "",name: "", co_partner: "",
      seller_name: "", check_number: "", originer_imprint: "", security: "", notification: ""
    };
  },
  methods: {
    onClick1() {
      this.security = !this.security;
    },
    onClick2() {
      this.notification = !this.notification;
    },
  },
  mounted() {
    axios
      .get("/accounts/" + this.id)
      .then(
        (response) => (
          console.log(response),
          (this.accounts = response.body),
          (this.logo = response.data[0].logo),
          (this.company_name = response.data[0].company_name),
          (this.address_company = response.data[0].address_company),
          (this.phone = response.data[0].phone),
          (this.name = response.data[0].name),
          (this.co_partner = response.data[0].co_partner),
          (this.seller_name = response.data[0].seller_name),
          (this.check_number = response.data[0].check_number),
          (this.originer_imprint = response.data[0].originer_imprint),
          (this.security = response.data[0].security),
          (this.notification = response.data[0].notification)
        )
      )
      .catch((error) => console.log(error));
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
/* Main */
.main-content {
  display: flex;
  margin-bottom: 80px;
}
/* Content Left */
.content-left {
  margin-right: 270px;
  flex: 6;
  display: flex;
  flex-direction: column;
}
.content-left-item {
  margin-bottom: 50px;
}
.content-left h4 {
  margin: 0;
  margin-bottom: 10px;
  color: #271b5a;
}
/* Infomation */
.line {
  border: 1px solid #201a39;
  margin: 0;
}
.infomation {
  display: grid;
  grid-template-columns: 20% 70% 10%;
  height: 80px;
  border-bottom: 1px solid #eeeeee;
}
.infomation > p {
  margin: auto 10px;
  color: #222222;
  font-size: 14px;
  font-weight: 500;
}
.infomation > label {
  margin: auto 20px;
  margin-left: 20px;
  color: #564c7e;
  font-size: 12px;
  font-weight: 500;
}
.image-info {
  width: 40px;
  height: 40px;
  margin: 20px 0 19px 10px;
}
.edit {
  display: grid;
  grid-template-columns: 50% 50%;
  width: 80px;
  border: none;
  background-color: white;
}
.edit > img {
  width: 20px;
  height: 20px;
  margin: auto;
}
.edit label {
  margin: auto;
  color: #d64d10;
}
.toggle {
  margin: auto 40px;
}
/* PopUp Image */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 98;
  background-color: #222222;
}
.edit-photo {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: #ffffff;
  width: 600px;
  height: 550px;
  border-radius: 10px;
  z-index: 99;
}
.title {
  background-color: white;
  height: 30px;
  width: 500px;
  margin: 50px;
  display: flex;
  justify-content: space-between;
}
.title h3 {
  margin: 0;
  color: #271b5a;
}
.title button {
  padding: 0;
  border: none;
  background-color: white;
}

.images {
  background-color: #f8f2e9;
  width: 495px;
  height: 300px;
  margin: 50px 50px 20px 50px;
  border: 2px dashed #d64d10;
  text-align: center;
}
.images img {
  margin-top: 50px;
  margin-bottom: 20px;
  width: 100px;
  height: 100px;
}
.upload-image {
  width: 100%;
  display: flex;
  justify-content: center;
  height: 20px;
}
.upload-image p {
  font-family: "NotoSansJP-Medium";
  font-weight: 700;
  font-size: 14px;
  color: #d64d10;
  margin: 0;
  margin-left: 5px;
}
.upload-image img {
  margin-right: 5px;
  margin-top: 0;
  width: 20px;
  height: 20px;
}
.update {
  border-radius: 5px;
  font-size: 14px;
  font-family: "NotoSansJP-Bold";
  width: 500px;
  height: 50px;
  color: white;
  background-color: #d64d10;
  margin: 0 50px 0 50px;
  cursor: pointer;
  border: none;
  font-weight: 700;
}
/* PopUp Name */
.edit-name {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: #ffffff;
  width: 600px;
  height: 320px;
  border-radius: 10px;
  z-index: 99;
}
.input-name {
  margin: 50px 50px 20px 50px;
}
.input-name label {
  color: #564c7e;
}
.input-name input {
  color: #222222;
  border: 1px solid #bbbbbb;
  border-radius: 5px;
  width: 100%;
  height: 50px;
  padding: 15px;
  box-sizing: border-box;
  font-weight: bold;
}
/* Add Content */
.add-content {
  display: flex;
  justify-content: space-between;
}
.add-content-submit {
  display: flex;
  border: none;
  background-color: white;
  color: #d64d10;
}
.add-content-submit img,
label {
  margin: auto;
}
/* Content Right */
.content-right {
  position: fixed;
  right: 30px;
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 190px;
}
.content-right-item {
  display: flex;
}
.content-right-item a {
  margin: auto 10px;
  font-size: 14px;
  font-weight: 500;
}
.content-right .active {
  color: #d64d10;
  background-color: #f8f2e9;
}
.content-right > div {
  color: #271b5a;
  width: 240px;
  height: 40px;
}
/* Footer */
.footer-content {
  margin: 50px auto;
  text-align: center;
}
.footer-content p {
  margin: 0;
  color: #666666;
  font-weight: 500;
}
</style>
