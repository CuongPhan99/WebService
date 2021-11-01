<template>
  <div class="main-content">
    <div class="content-left">
      <div class="list-search">
        <input type="text" value="連絡先を検索する" />
        <img src="../assets/images/icon_search.png" alt="" />
      </div>
      <div class="list-member">
        <table>
          <thead>
            <tr>
              <th>メンバー名</th>
              <th>所属部署/グループ</th>
              <th>会社名</th>
              <th>
                <input type="checkbox" id="cbx" style="display: none" />
                <label for="cbx" class="check">
                  <svg width="18px" height="18px" viewBox="0 0 18 18">
                    <path
                      d="M1,9 L1,3.5 C1,2 2,1 3.5,1 L14.5,1 C16,1 17,2 17,3.5 L17,14.5 C17,16 16,17 14.5,17 L3.5,17 C2,17 1,16 1,14.5 L1,9 Z"
                    ></path>
                    <polyline points="1 9 7 14 15 4"></polyline></svg
                  ><span>すべて表示</span>
                </label>
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(contact, index) in customers" :key="index">
              <td>
                <div class="item">
                  <div class="item1">
                    <img src="../assets/images/profile.png" />
                  </div>
                  <div class="item2">
                    <div class="item3">
                      <p>
                        <span>{{ contact.last_name }}</span>
                        <span>{{ contact.first_name }}</span>
                      </p>
                      <img src="../assets/images/icon_hide_black.png" alt="" />
                    </div>
                    <div class="item4">
                      <p>{{ contact.email }}</p>
                    </div>
                  </div>
                </div>
              </td>
              <td>所属部署</td>
              <td><button class="company">COMPANY A</button></td>
              <td>
                <div class="tooltip">
                  <button v-on:click="myFunction(index)">
                    <img src="../assets/images/Group 7264.png" />
                  </button>
                  <span v-show="showOption == index" class="tooltiptext tooltip-bottom">
                    <a @click="editContact = true"
                      ><img src="../assets/images/icon_edit.png" />編集</a
                    >
                    <a @click="hideContact = true"
                      ><img
                        src="../assets/images/icon_hide.png"
                      />非表示にする</a
                    >
                    <a @click="deleteContact = true"
                      ><img src="../assets/images/icon_delete.png" />削除</a
                    >
                  </span>
                </div>
                <div>
                  <transition name="fade" appear>
                    <div
                      class="modal-overlay"
                      v-if="editContact"
                      @click="editContact = false"
                    ></div>
                  </transition>
                  <transition name="slide" appear>
                    <div class="contact" v-if="editContact">
                      <div class="title" @click="editContact = false">
                        <h3>基本情報の編集</h3>
                        <button>
                          <img src="../assets/images/close.png" alt="" />
                        </button>
                      </div>
                      <div class="main-contact">
                        <div class="full-name">
                          <label><span>【必須】</span>氏名</label>
                          <input type="text" value="公的" />
                          <input type="text" value="太郎" />
                        </div>
                        <div class="email">
                          <label><span>【必須】</span>メールアドレス</label>
                          <input type="text" value="taro_yamada@gmail.com" />
                        </div>
                        <div class="company-name">
                          <label>（任意）企業名</label>
                          <input type="text" value="COMPANY A" />
                        </div>
                        <div class="address">
                          <label>（任意）所属部署/グループ </label>
                          <input type="text" value="所属部署" />
                        </div>
                        <button class="edit-submit">変更する</button>
                      </div>
                    </div>
                  </transition>
                </div>
                <div>
                  <transition name="fade" appear>
                    <div
                      class="modal-overlay"
                      v-if="hideContact"
                      @click="hideContact = false"
                    ></div>
                  </transition>
                  <transition name="slide" appear>
                    <div class="contact hide" v-if="hideContact">
                      <div class="title" @click="hideContact = false">
                        <h3>連絡先の非表示</h3>
                        <button>
                          <img src="../assets/images/close.png" alt="" />
                        </button>
                      </div>
                      <div class="hide-contact">
                        <div class="question">
                          <p>
                            <span>公的</span>
                            <span>太郎</span>を非表示にしますか？
                          </p>
                          <p>
                            連絡先は削除されず、一覧に表示されなくなります。
                          </p>
                        </div>
                        <button class="edit-submit">非表示にする</button>
                      </div>
                    </div>
                  </transition>
                </div>
                <div>
                  <transition name="fade" appear>
                    <div
                      class="modal-overlay"
                      v-if="deleteContact"
                      @click="deleteContact = false"
                    ></div>
                  </transition>
                  <transition name="slide" appear>
                    <div class="contact hide" v-if="deleteContact">
                      <div class="title" @click="deleteContact = false">
                        <h3>連絡先の削除</h3>
                        <button>
                          <img src="../assets/images/close.png" alt="" />
                        </button>
                      </div>
                      <div class="hide-contact">
                        <div class="question">
                          <p>
                            <span>公的</span>
                            <span>太郎</span>を削除してもよろしいですか？
                          </p>
                          <p>削除すると元に戻すことはできません。</p>
                        </div>
                        <button class="edit-submit">削除する</button>
                      </div>
                    </div>
                  </transition>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="content-right">
      <button class="update">
        <img src="../assets/images/Union 6.png" alt="" />
        <p>契約書を作成する</p>
      </button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      showOption: -1,
      editContact: false,
      hideContact: false,
      deleteContact: false,
      customers: [],
    };
  },
  methods: {
    myFunction(index) {
      this.showOption = index;
    },
  },
  mounted() {
    axios
      .get("/customers")
      .then((response) => (this.customers = response.data))
      .catch((error) => console.log(error));
  },
};
</script>

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
.list-search {
  position: fixed;
  width: calc(100% - 600px);
  box-shadow: 0px -18px 0px 13px white;
  z-index: 98;
}
.list-search > input {
  width: 100%;
  height: 50px;
  padding-left: 20px;
  border: 1px solid #dddddd;
  border-radius: 5px;
  background-color: #fafafa;
}
.list-search img {
  position: absolute;
  right: 15px;
  margin: 15px auto;
}
/* List Item */
.list-member {
  margin-top: 50px;
}
table {
  display: grid;
  border-collapse: collapse;
  min-width: 100%;
  grid-template-columns:
    minmax(150px, 1.67fr)
    minmax(150px, 1fr)
    minmax(150px, 1fr)
    minmax(150px, 1.67fr);
}
thead,
tbody,
tr {
  display: contents;
}
th,
td {
  display: grid;
  place-items: center;
  font-size: 12px;
  font-weight: 500;
  text-overflow: ellipsis;
  white-space: nowrap;
}
th {
  position: sticky;
  top: 220px;
  padding-top: 10px;
  background-color: white;
  color: #564c7e;
  border-bottom: 4px solid #201a39;
  z-index: 90;
}
th:first-child,
td:first-child {
  justify-content: left;
  padding-left: 20px;
}
th:last-child,
td:last-child {
  justify-content: right;
}
td {
  padding: 2px;
  border-bottom: 1px solid #dddddd;
}
td .company {
  background-color: white;
  border: 1px solid #dddddd;
  border-radius: 15px;
  width: 150px;
  height: 30px;
}
.item {
  display: flex;
  height: 40px;
}
.item1 {
  margin-right: 20px;
  height: 40px;
}
.item1 > img,
.item2 {
  height: 40px;
}
.item3 {
  display: grid;
  grid-template-columns: auto auto;
}
.item3 > p {
  margin: 0;
  text-align: left;
  font-size: 14px;
  font-weight: 700;
  color: #222222;
}
.item4 > p {
  margin: 0;
  text-align: left;
  width: 70px;
  color: #666666;
}

/* In put check all */
.check {
  display: flex;
  float: right;
}
.check > span {
  margin: 10px;
}
.check svg {
  position: relative;
  z-index: 1;
  fill: none;
  stroke-linecap: round;
  stroke-linejoin: round;
  stroke: #c8ccd4;
  stroke-width: 1.5;
  transform: translate3d(0, 0, 0);
  transition: all 0.2s ease;
  margin: 10px 0 10px 0;
}
.check svg path {
  stroke: #c8ccd4;
  stroke-dashoffset: 0;
}
.check svg polyline {
  stroke-dasharray: 22;
  stroke-dashoffset: 66;
}
#cbx:checked + .check svg {
  stroke: #d64d10;
}
#cbx:checked + .check svg polyline {
  stroke-dashoffset: 42;
  transition: all 0.2s linear;
  transition-delay: 0.15s;
}
/* Option Item */
.tooltip {
  position: relative;
  float: right;
  margin: 15px auto;
  width: 30px;
  height: 30px;
  line-height: 35px;
  display: block;
}
.tooltip button {
  padding: 0;
  color: white;
  font-size: 16px;
  width: 30px;
  height: 30px;
  border: none;
  border-radius: 5px;
  z-index: 89;
}
.tooltip .tooltiptext {
  display: grid;
  position: absolute;
  width: 250px;
  background-color: #ffffff;
  color: #fff;
  text-align: center;
  padding: 5px 0;
  border: 1px solid rgb(0, 0, 0, 0.2);
  border-radius: 6px;
  box-shadow: 0 -5px 3px -2px rgb(0, 0, 0, 0.2);
  z-index: 88;
}

.tooltip-bottom {
  top: 120%;
  right: 0;
}
.tooltip .tooltip-bottom::after {
  content: "";
  position: absolute;
  bottom: 100%;
  left: 94%;
  margin-left: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: transparent transparent #ffffff transparent;
}
.tooltip .tooltip-bottom a {
  display: flex;
  align-items: center;
  height: 50px;
  text-decoration: none;
  color: #d64d10;
  font-size: 14px;
  font-weight: 500;
}
.tooltip .tooltip-bottom img {
  margin: 20px;
}
.tooltip .tooltip-bottom a:nth-child(3) {
  color: #d00d41;
}
/* PopUp Edit Contact */
.main-contact {
  display: flex;
  margin: 50px;
}
.main-contact > div {
  margin-bottom: 20px;
  text-align: left;
  display: flex;
  flex-direction: column;
}
.main-contact .full-name {
  display: grid;
  grid-column-gap: 10px;
}
.main-contact label {
  color: #564c7e;
  grid-column: 1 / span 2;
}
.main-contact span {
  color: #d64d10;
}

.main-contact input {
  padding: 15px;
  height: 50px;
  color: #222222;
  font-size: 14px;
  font-weight: 500;
  border: 1px solid #bbbbbb;
  border-radius: 5px;
}
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #222222;
  z-index: 98;
}
.contact {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: #ffffff;
  width: 600px;
  height: 590px;
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
.main-contact {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.edit-submit {
  border-radius: 5px;
  font-size: 14px;
  font-family: "NotoSansJP-Bold";
  width: 500px;
  height: 50px;
  color: white;
  background-color: #d64d10;
  border: none;
  font-weight: 700;
}
/* PopUp Hide, Delete Contact */
.contact.hide {
  height: 328px;
}
.hide-contact {
  margin: 50px;
}
.hide-contact .question {
  margin: 50px;
  text-align: center;
}
.hide-contact .question p {
  margin: 10px;
  font-size: 14px;
  font-weight: 500;
  color: #666666;
}
.hide-contact .question span {
  margin-left: 10px;
  color: #564c7e;
}

/* PopUp Delete Contact */

/* Content Right */
.content-right {
  position: fixed;
  right: 0;
  flex: 1;
}
.update {
  border-radius: 5px;
  font-size: 14px;
  width: 240px;
  height: 50px;
  color: white;
  background-color: #d64d10;
  margin: 0px 30px 0px 30px;
  border: none;
  font-weight: 700;
  font-size: medium;
  position: relative;
}
.update p {
  position: absolute;
  margin: 12px auto;
  left: 50px;
  top: 0;
}
.update img {
  margin: 15px;
  float: left;
}
</style>
