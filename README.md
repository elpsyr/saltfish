# saltfish



## 界面
<picture>
  <img width="400" alt="Example of run salt fish" src="./images/examples/main.png">
</picture>

## 功能
- 隐藏/显示 
点击hide/show 隐藏/显示咸鱼之王pc端小程序应用
- 一键奖励
点击自动升级奖励，领取奖励
- 一键钓鱼
点击自动钓鱼
- 自动任务
每隔6小时自动领取挂机奖励，每隔8小时自动钓鱼
<picture>
  <img width="400" alt="Example of run salt fish" src="./images/examples/reward.png">
</picture>
<picture>
  <img width="400" alt="Example of run salt fish" src="./images/examples/fishing.png">
</picture>

## 构建
```shell
cd cmd/app/
fyne package -os windows -icon ../../images/fish.ico
```