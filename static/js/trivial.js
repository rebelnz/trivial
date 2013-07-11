$(function(){

    var btn = $("#get-wdata");
    var titleList = $("#title-list");
    var titleDiv = $("#title");
    var contentDiv = $("#content");
    
    btn.click(function() {
        $.get("/wdata", function(data) {			
            for (i = 0; i < data["query"]["random"].length; i++) {
                var uri = encodeURIComponent(data["query"]["random"][i]["title"]);
                titleList.prepend('<li><a href="/wdetail?' 
	         + uri
	         + '">'
	         + data["query"]["random"][i]["title"] 
	         + '</a></li>');
            }
        }, 'json');		    
    });	


    $('#title-list').on('click','a',function(e) {
        var uri = $(this).attr('href');
        $.get(uri, function(data) {
            for (var pageId in data.query.pages) {
                if (data.query.pages.hasOwnProperty(pageId)) {
                    var title = data.query.pages[pageId].title;
                    var content = parseWData(data.query.pages[pageId].revisions[0]['*']);
                }
            }

            titleDiv.html(title);
            contentDiv.html(content);
			
        }, 'json');		
        e.preventDefault();
    });

});

function parseWData(s) {
	var matchTitles = /\[\[([^\]]+?)\|\'\'(.+?)\'\']\]/ig;
	var matchLinks = /\[\[(.+?)\]\]/ig;
	var source = s.replace(matchTitles, '$2');
	source = source.replace(matchLinks, '<a href="http://en.wikipedia.org/wiki/$1">$1</a>');
	
	// ss = source.split("'''");
	// if (ss[0] !== "") {		
	// 	return ss[0];
	// } else {
	return source;
	//}	
}


