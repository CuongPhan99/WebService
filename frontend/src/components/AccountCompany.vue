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
              :src="image == '' ? account.logo : image"
              alt=""
            />
            <button class="edit" @click="showImage = true">
              <img src="../assets/images/edit.png" alt="" /><label>編集</label>
            </button>
            <!-- PopUp Edit Image -->
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showImage"
                @click="showImage = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-photo" v-if="showImage">
                <div class="title">
                  <h3>企業ロゴを変更</h3>
                  <button @click="showImage = false">
                    <img src="../assets/images/close.png" alt="" />
                  </button>
                </div>
                <div class="images">
                  <img :src="image == '' ? account.logo : image" />
                  <div class="upload-image">
                    <img src="../assets/images/icon_upload.png" alt="" />
                    <button v-on:click="handleClickInputFile">
                      画像のアップロード
                    </button>
                    <input
                      @change="onFileChange"
                      ref="fileInputLogo"
                      type="file"
                      style="display: none"
                    />
                  </div>
                </div>
                <button
                  v-on:click="
                    updateLogo();
                    showImage = false;
                  "
                  type="submit"
                  class="update"
                >
                  変更する
                </button>
              </div>
            </transition>
            <!-- ----------- -->
          </div>
          <div class="infomation">
            <label>企業名</label>
            <p>{{ account.company_name }}</p>
            <button class="edit" @click="showCompanyName = true">
              <img src="../assets/images/edit.png" alt="" /><label>編集</label>
            </button>
            <!-- PopUp Edit Company Name -->
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showCompanyName"
                @click="showCompanyName = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-name" v-if="showCompanyName">
                <div class="title">
                  <h3>企業名の変更</h3>
                  <button @click="showCompanyName = false">
                    <img src="../assets/images/close.png" alt="" />
                  </button>
                </div>
                <form
                  @submit.prevent="updateCompanyName()"
                  action=""
                  method="post"
                >
                  <div class="input-name">
                    <label for="">企業名</label>
                    <input
                      id="companyName"
                      v-model="account.company_name"
                      name="companyName"
                      type="text"
                    />
                  </div>
                  <button
                    type="submit"
                    class="update"
                    @click="showCompanyName = false"
                  >
                    変更する
                  </button>
                </form>
              </div>
            </transition>
            <!-- ------------ -->
          </div>
          <div class="infomation">
            <label>企業住所</label>
            <p>{{ account.address_company }}</p>
            <button class="edit" @click="showAddressCompany = true">
              <img src="../assets/images/edit.png" alt="" /><label>編集</label>
            </button>
            <!-- PopUp Edit Address Company -->
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showAddressCompany"
                @click="showAddressCompany = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-name" v-if="showAddressCompany">
                <div class="title">
                  <h3>企業名の変更</h3>
                  <button @click="showAddressCompany = false">
                    <img src="../assets/images/close.png" alt="" />
                  </button>
                </div>
                <form
                  @submit.prevent="updateAddressCompany()"
                  action=""
                  method="post"
                >
                  <div class="input-name">
                    <label for="">企業住所</label>
                    <input
                      id="companyName"
                      v-model="account.address_company"
                      name="companyName"
                      type="text"
                    />
                  </div>
                  <button
                    type="submit"
                    class="update"
                    @click="showAddressCompany = false"
                  >
                    変更する
                  </button>
                </form>
              </div>
            </transition>
            <!-- ------------ -->
          </div>
          <div class="infomation">
            <label>電話番号</label>
            <p>{{ account.phone }}</p>
            <button class="edit" @click="showPhone = true">
              <img src="../assets/images/edit.png" alt="" /><label>編集</label>
            </button>
            <!-- PopUp Edit Phone -->
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showPhone"
                @click="showPhone = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-name" v-if="showPhone">
                <div class="title">
                  <h3>企業名の変更</h3>
                  <button @click="showPhone = false">
                    <img src="../assets/images/close.png" alt="" />
                  </button>
                </div>
                <form
                  @submit.prevent="updatePhone()"
                  action=""
                  method="post"
                >
                  <div class="input-name">
                    <label for="">電話番号</label>
                    <input
                      id="companyName"
                      v-model="account.phone"
                      name="companyName"
                      type="text"
                    />
                  </div>
                  <button
                    type="submit"
                    class="update"
                    @click="showPhone = false"
                  >
                    変更する
                  </button>
                </form>
              </div>
            </transition>
            <!-- ------------ -->
          </div>
          <div class="infomation">
            <label>代表者氏名</label>
            <p>{{ account.name }}</p>
            <button class="edit" @click="showName = true">
              <img src="../assets/images/edit.png" alt="" /><label>編集</label>
            </button>
            <!-- PopUp Edit Name -->
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showName"
                @click="showName = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-name" v-if="showName">
                <div class="title">
                  <h3>企業名の変更</h3>
                  <button @click="showName = false">
                    <img src="../assets/images/close.png" alt="" />
                  </button>
                </div>
                <form
                  @submit.prevent="updateName()"
                  action=""
                  method="post"
                >
                  <div class="input-name">
                    <label for="">代表者氏名</label>
                    <input
                      id="companyName"
                      v-model="account.name"
                      name="companyName"
                      type="text"
                    />
                  </div>
                  <button
                    type="submit"
                    class="update"
                    @click="showName = false"
                  >
                    変更する
                  </button>
                </form>
              </div>
            </transition>
            <!-- ------------ -->
          </div>
          <div class="infomation">
            <label>パートナー会社</label>
            <p>{{ account.co_partner }}</p>
          </div>
          <div class="infomation">
            <label>紹介営業担当者</label>
            <p>{{ account.seller_name }}</p>
          </div>
          <div class="infomation">
            <label>管理番号</label>
            <p>{{ account.check_number }}</p>
          </div>
          <div class="infomation">
            <label>オリジナル印影</label>
            <img
              class="image-info"
              :src="
                originer_imprint == ''
                  ? account.originer_imprint
                  : originer_imprint
              "
              alt=""
            />
            <button class="edit" @click="showOriginerImprint = true">
              <img src="../assets/images/edit.png" alt="" /><label>編集</label>
            </button>
            <!-- PopUp Edit Image Originer Imprint -->
            <transition name="fade" appear>
              <div
                class="modal-overlay"
                v-if="showOriginerImprint"
                @click="showOriginerImprint = false"
              ></div>
            </transition>
            <transition name="slide" appear>
              <div class="edit-photo" v-if="showOriginerImprint">
                <div class="title">
                  <h3>企業ロゴを変更</h3>
                  <button @click="showOriginerImprint = false">
                    <img src="../assets/images/close.png" alt="" />
                  </button>
                </div>
                <div class="images">
                  <img
                    :src="
                      originer_imprint == ''
                        ? account.originer_imprint
                        : originer_imprint
                    "
                  />
                  <div class="upload-image">
                    <img src="../assets/images/icon_upload.png" alt="" />
                    <button v-on:click="handleClickInputFile">
                      画像のアップロード
                    </button>
                    <input
                      @change="onFileOriginerImprintChange"
                      ref="fileInputLogo"
                      type="file"
                      style="display: none"
                    />
                  </div>
                </div>
                <button
                  v-on:click="
                    updateOriginerImprint();
                    showOriginerImprint = false;
                  "
                  type="submit"
                  class="update"
                >
                  変更する
                </button>
              </div>
            </transition>
            <!-- ----------- -->
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
                  ? require('../assets/images/toggle-on.png')
                  : require('../assets/images/toggle-off.png')
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
                  ? require('../assets/images/toggle-on.png')
                  : require('../assets/images/toggle-off.png')
              "
              alt=""
            />
          </div>
        </div>
      </div>
      <div class="add-content">
        <h4>アラート通知先の設定</h4>
        <button class="add-content-submit">
          <img src="../assets/images/icon_add.png" alt="" /><label
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
import axios from "axios";

export default {
  data() {
    return {
      showImage: false,
      showCompanyName: false,
      showAddressCompany: false,
      showPhone: false,
      showName: false,
      showOriginerImprint: false,
      id: this.$route.params.id,
      account: [],
      security: "",
      notification: "",
      image: "",
      originer_imprint: "",
    };
  },
  methods: {
    onClick1() {
      this.security = !this.security;
    },
    onClick2() {
      this.notification = !this.notification;
    },
    handleClickInputFile() {
      this.$refs.fileInputLogo.click();
    },
    onFileChange(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) return;
      this.createImage(files[0]);
    },
    createImage(file) {
      var reader = new FileReader();
      reader.onload = (e) => {
        this.image = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    onFileOriginerImprintChange(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) return;
      this.createImage2(files[0]);
    },
    createImage2(file) {
      var reader = new FileReader();
      reader.onload = (e) => {
        this.originer_imprint = e.target.result;
      };
      reader.readAsDataURL(file);
    },
    updateLogo() {
      const fd = new FormData();
      fd.append("logo", this.image);
      axios.post("/accounts/" + this.id, fd);
      this.$forceUpdate();
    },
    updateCompanyName() {
      const fd = new FormData();
      fd.append("company_name", this.account.company_name);
      axios.post("/accounts/" + this.id, fd);
    },
    updateAddressCompany() {
      const fd = new FormData();
      fd.append("address_company", this.account.address_company);
      axios.post("/accounts/" + this.id, fd);
    },
    updatePhone() {
      const fd = new FormData();
      fd.append("phone", this.account.phone);
      axios.post("/accounts/" + this.id, fd);
    },
    updateName() {
      const fd = new FormData();
      fd.append("name", this.account.name);
      axios.post("/accounts/" + this.id, fd);
    },
    updateOriginerImprint() {
      const fd = new FormData();
      fd.append("originer_imprint", this.originer_imprint);
      axios.post("/accounts/" + this.id, fd);
    },
  },
  mounted() {
    axios
      .get("/accounts/" + this.id)
      .then(
        (response) => (
          (this.account = response.data[0]),
          (this.security = this.account.security),
          (this.notification = this.account.notification)
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
.upload-image button {
  font-family: "NotoSansJP-Medium";
  font-weight: 700;
  font-size: 14px;
  color: #d64d10;
  margin: 0;
  margin-left: 5px;
  border: none;
  background-color: transparent;
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
