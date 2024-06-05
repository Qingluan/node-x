data => {
    let Xpath = (path, ele=null, filter=null)=>{
        if (ele == null){
            ele = document;
        }
        var result = document.evaluate(path, ele, null, XPathResult.ANY_TYPE, null);
        var nodes = [];
        var node = result.iterateNext();
        // 遍历所有匹配的节点
        while (node) {
            // 检查父节点是否是script或style标签
            var parent = node.parentNode;
            if (parent && parent.tagName.toLowerCase() !== 'script' && parent.tagName.toLowerCase() !== 'style') {
                // 如果不是，则将文本节点的内容添加到数组中
                if (node.textContent.trim().length > 0){
                    if (filter != null && typeof(filter) == "function"){
                        if (filter(node)){
                            nodes.push(node);
                        }
                    }else{
                        nodes.push(node);
                    }
                }else{
                }
            }
            node = result.iterateNext();
        }
        return nodes
    }

    let eles = Xpath("//a[@data-ved]/div/div")
    let searchItems = [];
    for(let div of eles){
        if (div.children.length > 2 && div.children.length < 5){
            let post = div.children[0].textContent.trim();
            let title = div.children[1].textContent.trim();
            let url = div.parentElement.parentElement.href;
            let date = div.children[2].textContent.trim();    
            let item = {
                title: title,
                url: url,
                post: post,
                date: date,
            }
            searchItems.push(item);
            
        }else if (div.children.length > 4){
            let url = div.parentElement.parentElement.href;
            let post = div.children[0].textContent.trim();
            let title = div.children[1].textContent.trim();
            let desc = div.children[2].textContent.trim();
            let date = div.children[4].textContent.trim();
            let item = {
                title: title,
                url: url,
                desc: desc,
                date: date,
                post: post,
            }
            searchItems.push(item);
        }
        
        
    }
    return JSON.stringify(searchItems);
}