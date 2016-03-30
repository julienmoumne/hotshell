var Item = React.createClass({
  render: function() {
    return <div className="item">
        <span className='key'>{this.props.itemKey}</span>
        <span className={this.props.itemKey == '' || this.props.item.desc == '' ? '' : 'item-desc'} dangerouslySetInnerHTML={{__html: this.props.item.desc}}></span>
        <span className='cmd'>{this.props.item.cmd == '' && this.props.item.items.length > 0 ? '>' : this.props.item.cmd}</span>
    </div>;
  }
});

var Menu = React.createClass({
  render: function() {
    var items = this.props.item.items.map(function(el) {
      	return <Item itemKey={el.key} item={el.item}/>
    })

    var trailLength = this.props.trail.length
    var trail = this.props.trail.map(function(item, ix){
        var isActive = ix == trailLength - 1
        return <div className={isActive ? 'active-menu' : 'parent-menu'} key={item.key}>{item.desc}</div>
    })

    return (
      <div className="menu">
        {trail}
        <div className="items">{items}</div>
        <div className="actions"><span className="key">spacebar</span> back</div>
        <div className="prompt">
            ? {this.props.userChoice ? this.props.userChoice : <span className='blinking-cursor'>|</span>}
        </div>
      </div>
    );
  }
});

var MenuController = React.createClass({
  getInitialState: function() {
    return {activatedItems: [{item: this.props.item, trail: [this.props.item]}]};
  },
  handleUserChoice: function(userChoice) {

    if (userChoice == '') return

    var lastActivatedItem = this.state.activatedItems[this.state.activatedItems.length - 1]

    var trail = lastActivatedItem.trail.slice()
    var nextItem
    if (userChoice == ' ') {
        if (trail.length == 1) return
        trail.pop()
        nextItem = trail[trail.length - 1]
        userChoice = 'spacebar'
    } else {
        nextItem = _.find(lastActivatedItem.item.items, function(el) {
            return el.key == userChoice
        })
        if (nextItem == undefined) return
        nextItem = nextItem.item
        if (nextItem.cmd != '') {
            this.state.activatedItems.push({item: nextItem})
            nextItem = lastActivatedItem.item
        } else {
            trail.push(nextItem)
        }
    }

    lastActivatedItem.userChoice = userChoice

    this.state.activatedItems.push({item: nextItem, trail: trail})
    this.setState({activatedItems: this.state.activatedItems})

    window.scrollTo(0,document.body.scrollHeight);
  },
  componentDidMount(){
      var self = this
      document.body.onkeypress = function(e){
        self.handleUserChoice(String.fromCharCode(e.which))
      };
  },
  render: function() {

    var items = this.state.activatedItems.map(function(el) {
      	return el.item.cmd == '' ?
      	    <Menu trail={el.trail} item={el.item} userChoice={el.userChoice}/> :
      	    <div>
      	        <div className='executed-cmd'>/bin/bash -c '{el.item.cmd}'</div>
      	        <div className='stdout'>
      	            (the output of the command would be displayed here)
      	        </div>
            </div>
    })
    var time = moment().format('HH:mm');
    return (
      <div className="menuController">
        <div><span className="time">{time}</span> $ hs -f ./{filename}</div>
        <br/>
        {items}
      </div>
    );
  }
});

ReactDOM.render(
  <MenuController item={item} />,
  document.getElementById('container')
);